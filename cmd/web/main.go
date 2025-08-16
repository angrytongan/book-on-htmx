package main

import (
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"
)

const (
	port = 8880
)

func run() error {
	app, err := newApplication()
	if err != nil {
		return fmt.Errorf("newApplication(): %w", err)
	}

	// Navigation.
	app.mux.Get("/nav/", app.nav)
	app.mux.Get("/nav/{href}", app.nav)

	// Pages.
	app.mux.Get("/", app.home)
	app.mux.Get("/tabs", app.tabs)
	app.mux.Get("/theme", app.theme)
	app.mux.Get("/search", app.search)

	app.mux.Group(func(r chi.Router) {
		r.Use(delayResponse(500))

		// Widgets.
		r.Get("/theme", app.theme)
		r.Post("/theme-chooser", app.themeChooserSave)

		r.Get("/tabs/links", app.tabsLinks)
		r.Get("/tabs/links/{tab}", app.tabsLinks)

		r.Get("/tabs/links-oob", app.tabsLinksOOB)
		r.Get("/tabs/links-oob/{tab}", app.tabsLinksOOB)
		r.Get("/tabs/buttons-oob", app.tabsButtonsOOB)
		r.Get("/tabs/buttons-oob/{tab}", app.tabsButtonsOOB)
		r.Get("/tabs/radio-oob", app.tabsRadiosOOB)
		r.Get("/tabs/radio-oob/{tab}", app.tabsRadiosOOB)
		r.Get("/tabs/content/{content}", app.tabsContent)

		r.Get("/search/term", app.searchTerm)
	})

	return app.listen(port)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
