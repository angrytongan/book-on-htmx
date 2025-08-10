package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"bonh/internal/nav"
	"bonh/internal/theme"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/mattn/go-sqlite3"
)

const (
	muxCompressionLevel = 5
)

type Application struct {
	mux *chi.Mux
	tpl *template.Template

	themeRepo theme.Repository
}

func newApplication() (*Application, error) {
	assetFileServer := http.FileServer(http.Dir("./assets"))
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Compress(muxCompressionLevel, "text/html", "text/css", "text/javascript"))
	mux.Handle("/css/*", assetFileServer)
	mux.Handle("/js/*", assetFileServer)

	tpl := template.Must(template.ParseGlob("templates/*.tmpl"))

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, fmt.Errorf("sql.Open(): %w", err)
	}

	themeRepo := theme.SQLRepository(db)

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

		activeTheme, err := app.themeRepo.Active(1)
		if err != nil {
			activeTheme = theme.DefaultTheme()
		}

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
) {
	log.Printf("%s %s: %v\n", r.Method, r.URL, err)
	http.Error(w, http.StatusText(statusCode), statusCode)
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
