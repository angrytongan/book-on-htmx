package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"bonh/internal/nav"
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

	themeRepo theme.Repository
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
		return nil, fmt.Errorf("pgxpool.New(%s): %w", mustGetenv("DATABASE_CONNECTION_STRING"), err)
	}

	themeRepo := theme.NewPGXPoolRepository(pool)

	return &Application{
		mux,
		tpl,
		themeRepo,
	}, nil
}

func (app *Application) render(w http.ResponseWriter,
	r *http.Request,
	block string,
	pageData map[string]any,
	statusCode int,
) {
	var b bytes.Buffer

	// Render a full page if we didn't get a htmx request.
	if r.Header.Get("Hx-Request") != "true" {
		// Setup non-specific page template data here.
		if pageData == nil {
			pageData = map[string]any{}
		}

		activeTheme, _ := app.themeRepo.Active(context.Background(), 1)

		pageData["DataTheme"] = activeTheme
		pageData["Nav"] = nav.PageLinks(r.URL.String())

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
		clientMessage = strings.Join(msg, " - ") // :shrug:
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
