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
	app.mux.Get("/settings", app.settings)

	// Widgets.
	app.mux.Get("/theme-chooser", app.themeChooser)
	app.mux.Post("/theme-chooser", app.themeChooserSave)

	app.mux.Get("/tabs/buttons", app.tabsButtons)
	app.mux.Get("/tabs/buttons/{tab}", app.tabsButtons)
	app.mux.Get("/tabs/radio-buttons", app.tabsRadioButtons)
	app.mux.Get("/tabs/radio-buttons/{tab}", app.tabsRadioButtons)
	app.mux.Get("/tabs/content/{content}", app.tabsContent)

	return app.listen(port)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
