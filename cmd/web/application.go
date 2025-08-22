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

// includeExtraBlockData adds data outside of that required for rendering the
// page to be added for inclusion when executing the template. Templates may
// rely on data that isn't native to themselves; for example, an endpoint that
// renders itself as a "full page" may require updating a navigation component
// showing the current page, but the endpoint doesn't (and shouldn't) know
// anything about the navigation component. We do that resolution here, keeping
// the endpoint (relatively) clean.
func (app *Application) includeExtraBlockData(
	r *http.Request,
	blockData map[string]any,
	extraBlockData []string,
) map[string]any {
	if blockData == nil {
		blockData = map[string]any{}
	}

	for _, extra := range extraBlockData {
		switch extra {
		case "navigation":
			blockData["Nav"] = nav.PageLinks(r.URL.Path)

		case "theme":
			activeTheme, _ := app.themeRepo.Active(context.Background(), 1)
			blockData["PageTheme"] = activeTheme
		}
	}

	return blockData
}

// render pulls together the data required for the template to be rendered, and
// executes it. statusCode is sent as the response.
func (app *Application) render(
	w http.ResponseWriter,
	r *http.Request,
	block string,
	blockData map[string]any,
	statusCode int,
	extraBlockData ...string,
) {
	var b bytes.Buffer

	// If the client is doing a full page load, add in any "global" data for
	// the full page template block(s). This will typically be stuff that is
	// present in the page header or footer, eg. theme name for the data-theme
	// attribute in <html>.
	if r.Header.Get("Hx-Request") != "true" {
		extraBlockData = append(extraBlockData, "theme", "navigation")

		// Full page loads require us to load the entire page. Blocks that are
		// for entire pages should have "-page" appended to them in the
		// template files. This allows each endpoint to specify it's own layout
		// in the "-page" block.
		block += "-page"
	}

	// Amend the block data to include anything that has been specified as
	// extra. We rely on the block template that is being rendered to contain
	// correct extra templates so this data can be shown.
	blockData = app.includeExtraBlockData(r, blockData, extraBlockData)

	// Run the template.
	err := app.tpl.ExecuteTemplate(&b, block, blockData)
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

	// Send to client.
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
