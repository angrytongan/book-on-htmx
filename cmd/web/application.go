package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"bonh/internal/dog"
	"bonh/internal/nav"
	"bonh/internal/search"
	"bonh/internal/theme"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	muxCompressionLevel = 5
)

var (
	ErrMissingPageTemplate = errors.New("missing page template")
)

type Application struct {
	mux *chi.Mux
	tpl *template.Template

	themeRepo  theme.Repository
	searchRepo search.Repository
	dogRepo    dog.Repository
	words      []string
}

func mustGetenv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("missing: %s", key)
	}

	return value
}

func loadWords() ([]string, error) {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return nil, fmt.Errorf("failed to open dictionary file: %w", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Printf("error closing dictionary file: %v", closeErr)
		}
	}()

	var words []string
	scanner := bufio.NewScanner(file)
	count := 0
	
	for scanner.Scan() && count < 200 {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read dictionary file: %w", err)
	}

	return words, nil
}

func newApplication() (*Application, error) {
	assetFileServer := http.FileServer(http.Dir("./assets"))
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Compress(muxCompressionLevel, "text/html", "text/css", "text/javascript"))
	mux.Handle("/css/*", assetFileServer)
	mux.Handle("/js/*", assetFileServer)

	funcMap := template.FuncMap{
		"title": func(s string) string {
			c := cases.Title(language.Und)
			return c.String(s)
		},
	}

	tpl := template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.tmpl"))

	pool, err := pgxpool.New(context.Background(), mustGetenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New(): %w", err)
	}

	themeRepo := theme.NewPGXPoolRepository(pool)
	searchRepo := search.NewPGXPoolRepository(pool)
	dogRepo := dog.NewPGXPoolRepository(pool)

	words, err := loadWords()
	if err != nil {
		return nil, fmt.Errorf("loadWords(): %w", err)
	}

	return &Application{
		mux,
		tpl,
		themeRepo,
		searchRepo,
		dogRepo,
		words,
	}, nil
}

// render pulls together the data required for the template to be rendered, and
// executes it. statusCode is sent as the response.
func (app *Application) render(
	w http.ResponseWriter,
	r *http.Request,
	block string,
	blockData map[string]any,
	statusCode int,
) {
	var b bytes.Buffer
	var templatesToRender []string

	if blockData == nil {
		blockData = map[string]any{}
	}

	pageBlock := "page-" + block

	if r.Header.Get("Hx-Request") != "true" {
		// This isn't a partial render, so we have to render the entire page. Grab
		// all the bits that are required to render an entire pgae, and construct
		// it from the bits.
		activeTheme, _ := app.themeRepo.Active(context.Background(), 1)
		blockData["PageTheme"] = activeTheme
		blockData["Nav"] = nav.PageLinks(r.URL.Path)

		// We must have a page block.
		if app.tpl.Lookup(pageBlock) == nil {
			app.serverError(
				w,
				r,
				fmt.Errorf("app.tpl.Lookup(%s): %w", pageBlock, ErrMissingPageTemplate),
				http.StatusInternalServerError,
			)

			return
		}

		templatesToRender = append(templatesToRender, pageBlock)
	} else {
		// We are doing a partial render.

		// If we have a page block for this partial, then load any bits that
		// may need to be rendered, such as navigation. Use the oob versions,
		// as these will be rendered out of band.
		if app.tpl.Lookup(pageBlock) != nil {
			blockData["Nav"] = nav.PageLinks(r.URL.Path)
			templatesToRender = append(templatesToRender, "nav-oob")
		}

		// Add in the requested block.
		templatesToRender = append(templatesToRender, block)
	}

	// Execute all the templates in turn and store in the output buffer.
	for _, t := range templatesToRender {
		if app.tpl.Lookup(t) == nil {
			app.serverError(
				w,
				r,
				fmt.Errorf("app.tpl.Lookup(%s): no such template", block),
				http.StatusInternalServerError,
			)
			return
		}

		var lb bytes.Buffer
		if err := app.tpl.ExecuteTemplate(&lb, t, blockData); err != nil {
			app.serverError(
				w,
				r,
				fmt.Errorf("app.tpl.ExecuteTemplate(%s): %w", t, err),
				http.StatusInternalServerError,
			)
			return
		}

		b.Write(lb.Bytes())
	}

	// NOTE hardcoded content type here
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)

	// Send final buffer to the client.
	if _, err := w.Write(b.Bytes()); err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("w.Write(): %w", err),
			http.StatusInternalServerError,
		)

		return
	}
}

func (app *Application) listen(port int) error {
	log.Printf("listening on port %d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.mux)

	return fmt.Errorf("http.ListenAndServe(%d): %w", port, err)
}

func (app *Application) serverError(
	w http.ResponseWriter,
	r *http.Request,
	err error,
	statusCode int,
	msg ...string,
) {
	log.Printf("%s %s: %v\n", r.Method, r.URL, err)

	clientMessage := http.StatusText(statusCode)
	if len(msg) > 0 {
		clientMessage = strings.Join([]string{clientMessage, strings.Join(msg, ":")}, " - ") // :shrug:
	}

	http.Error(w, clientMessage, statusCode)
}

func (app *Application) clientRedirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	if r.Header.Get("Hx-Request") == "true" {
		w.Header().Set("Hx-Redirect", url)
	} else {
		http.Redirect(w, r, url, code)
	}
}

func (app *Application) clientLocation(w http.ResponseWriter, r *http.Request, url string, code int) {
	if r.Header.Get("Hx-Request") == "true" {
		locationVars := map[string]string{
			"path":   url,
			"target": "#main",
		}

		bytes, err := json.Marshal(locationVars)
		if err != nil {
			http.Redirect(w, r, url, code)
		}

		w.Header().Set("HX-Location", string(bytes))
	} else {
		http.Redirect(w, r, url, code)
	}
}
