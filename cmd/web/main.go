package main

import (
	"fmt"
	"log"
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

	// Widgets.
	app.mux.Get("/theme", app.theme)
	app.mux.Post("/theme-chooser", app.themeChooserSave)

	app.mux.Get("/tabs/links", app.tabsLinks)
	app.mux.Get("/tabs/links/{tab}", app.tabsLinks)

	app.mux.Get("/tabs/links-oob", app.tabsLinksOOB)
	app.mux.Get("/tabs/links-oob/{tab}", app.tabsLinksOOB)
	app.mux.Get("/tabs/buttons-oob", app.tabsButtonsOOB)
	app.mux.Get("/tabs/buttons-oob/{tab}", app.tabsButtonsOOB)
	app.mux.Get("/tabs/radio-oob", app.tabsRadiosOOB)
	app.mux.Get("/tabs/radio-oob/{tab}", app.tabsRadiosOOB)
	app.mux.Get("/tabs/content/{content}", app.tabsContent)

	return app.listen(port)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
