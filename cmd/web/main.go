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

	app.mux.Post("/settings/theme", app.settingsTheme)

	return app.listen(port)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
