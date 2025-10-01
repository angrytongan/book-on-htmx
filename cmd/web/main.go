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

	// Pages.
	app.mux.Get("/", app.home)
	app.mux.Get("/dashboard", app.dashboard)
	app.mux.Get("/dog", app.dog)
	app.mux.Get("/drawer", app.drawer)
	app.mux.Get("/form", app.form)
	app.mux.Get("/leaflet", app.leaflet)
	app.mux.Get("/loading", app.loading)
	app.mux.Get("/map-form", app.mapForm)
	app.mux.Get("/repl", app.repl)
	app.mux.Get("/search", app.search)
	app.mux.Get("/tabs", app.tabs)
	app.mux.Get("/theme", app.theme)
	app.mux.Get("/toast", app.toast)
	app.mux.Get("/settings", app.settings)

	// htmx requests.
	app.mux.Group(func(r chi.Router) {
		r.Use(delayResponse(500))
		r.Use(app.mustBeHtmxRequest)

		// Widgets.
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

		r.Get("/toast/server-time", app.toastServerTime)
		r.Get("/toast/random-number", app.toastRandomNumber)
		r.Get("/toast/random-letter", app.toastRandomLetter)
		r.Get("/toast/random-word", app.toastRandomWord)

		r.Get("/weather", app.weather)
		r.Get("/weather/current", app.weatherCurrent)

		r.Get("/time", app.time)
		r.Get("/time/servertime", app.servertime)

		r.Get("/dog/table", app.dogTable)

		r.Get("/drawer/content", app.drawerContent)

		r.Get("/form/step1", app.formStep1)
		r.Post("/form/step1", app.formStep1Process)
		r.Get("/form/step2", app.formStep2)

		r.Get("/loading/thing/{number}", app.loadingThing)

		r.Get("/leaflet/location/{name}", app.leafletLocation)

		r.Get("/map-form/form", app.mapFormForm)
	})

	// api requests.
	app.mux.Group(func(r chi.Router) {
		r.Get("/api/v1/location/{name}", app.leafletLocationJSON)
	})

	return app.listen(port)
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
