package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrMissingField = errors.New("missing field")
)

type Location struct {
	Name string
	Lat  float64
	Lng  float64
}

var (
	locations = map[string]Location{
		"Sydney":    {"Sydney", -33.8727, 151.2057},
		"Melbourne": {"Melbourne", -37.8136, 144.9631},
		"Brisbane":  {"Brisbane", -27.4705, 153.0260},
		"Perth":     {"Perth", -31.9514, 115.8617},
		"Hobart":    {"Hobart", -42.8826, 147.3257},
		"Adelaide":  {"Adelaide", -34.9285, 138.6007},
		"Darwin":    {"Darwin", -12.4637, 130.8444},
	}
)

func (app *Application) leaflet(w http.ResponseWriter, r *http.Request) {
	blockData := map[string]any{
		"Locations": locations,
	}

	app.render(w, r, "leaflet", blockData, http.StatusOK)
}

func (app *Application) leafletLocation(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		app.serverError(
			w,
			r,
			fmt.Errorf("r.PathValue(name): %w", ErrMissingField),
			http.StatusBadRequest,
		)

		return
	}

	location, ok := locations[name]
	if !ok {
		app.serverError(
			w,
			r,
			fmt.Errorf("location(%s): %w", name, ErrMissingField),
			http.StatusBadRequest,
		)

		return
	}

	blockData := map[string]any{
		"Name": location.Name,
		"Lat":  location.Lat,
		"Lng":  location.Lng,
	}

	w.Header().Set("Hx-Trigger", `{"selected": {"url": "/api/v1/location/`+location.Name+`"}}`)
	app.render(w, r, "leaflet-location-info", blockData, http.StatusOK)
}

func (app *Application) leafletLocationJSON(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		app.serverError(
			w,
			r,
			fmt.Errorf("r.PathValue(name): %w", ErrMissingField),
			http.StatusBadRequest,
		)

		return
	}

	location, ok := locations[name]
	if !ok {
		app.serverError(
			w,
			r,
			fmt.Errorf("location(%s): %w", name, ErrMissingField),
			http.StatusBadRequest,
		)

		return
	}

	blockData := map[string]any{
		"json": map[string]any{
			"name": location.Name,
			"lat":  location.Lat,
			"lng":  location.Lng,
		},
	}

	app.render(w, r, "json", blockData, http.StatusOK)
}
