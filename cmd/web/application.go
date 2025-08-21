package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"bonh/internal/nav"
	"bonh/internal/search"
	"bonh/internal/theme"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	muxCompressionLevel = 5
)

type Application struct {
	mux *chi.Mux
	tpl *template.Template

	themeRepo  theme.Repository
	searchRepo search.Repository
}

func mustGetenv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("missing: %s", key)
	}

	return value
}

func newApplication() (*Application, error) {
	assetFileServer := http.FileServer(http.Dir("./assets"))
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Compress(muxCompressionLevel, "text/html", "text/css", "text/javascript"))
	mux.Handle("/css/*", assetFileServer)
	mux.Handle("/js/*", assetFileServer)

	tpl := template.Must(template.ParseGlob("templates/*.tmpl"))

	pool, err := pgxpool.New(context.Background(), mustGetenv("DATABASE_CONNECTION_STRING"))
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New(): %w", err)
	}

	themeRepo := theme.NewPGXPoolRepository(pool)
	searchRepo := search.NewPGXPoolRepository(pool)

	return &Application{
		mux,
		tpl,
		themeRepo,
		searchRepo,
	}, nil
}

func (app *Application) render(
	w http.ResponseWriter,
	r *http.Request,
	block string,
	pageData map[string]any,
	statusCode int,
) {
	var b bytes.Buffer

	if pageData == nil {
		pageData = map[string]any{}
	}

	// Setup anything required for full page load.
	if r.Header.Get("Hx-Request") != "true" {
		// Setup non-specific page template data here.

		activeTheme, _ := app.themeRepo.Active(context.Background(), 1)

		pageLoad := map[string]any{
			"DataTheme": activeTheme,
		}

		pageData["PageLoad"] = pageLoad

		block += "-page"
	}

	err := app.tpl.ExecuteTemplate(&b, block, pageData)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("app.tpl.ExecuteTemplate(%s): %w", block, err),
			http.StatusInternalServerError,
		)

		return
	}

	// NOTE hardcoded content type here
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)

	_, err = w.Write(b.Bytes())
	if err != nil {
		app.serverError(w, r, fmt.Errorf("w.Write(): %w", err), http.StatusInternalServerError)

		return
	}
}

func (app *Application) renderWithNav(
	w http.ResponseWriter,
	r *http.Request,
	block string,
	pageData map[string]any,
	statusCode int,
) {
	if pageData == nil {
		pageData = map[string]any{}
	}

	pageData["Nav"] = nav.PageLinks(r.URL.Path)
	app.render(w, r, block, pageData, statusCode)
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

/*
func (app *Application) delayResponse(ms int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Duration(ms) * time.Millisecond)
			next.ServeHTTP(w, r)
		})
	}
}
*/
