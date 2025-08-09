package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	mux *chi.Mux
	tpl *template.Template
}

func newApplication() *Application {
	assetFileServer := http.FileServer(http.Dir("./assets"))
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Compress(5, "text/html", "text/css", "text/javascript"))
	mux.Handle("/css/*", assetFileServer)
	mux.Handle("/js/*", assetFileServer)

	tpl := template.Must(template.ParseGlob("templates/*.tmpl"))

	return &Application{
		mux,
		tpl,
	}
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

func (app *Application) listen(port int) {
	log.Printf("listening on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), app.mux)
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
