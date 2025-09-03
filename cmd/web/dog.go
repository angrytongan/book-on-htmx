package main

import (
	"fmt"
	"net/http"
)

type Dog struct {
	Colour string
	Breed  string
	Name   string
}

func (app *Application) dog(w http.ResponseWriter, r *http.Request) {
	blockData := map[string]any{
		"QP": r.URL.Query().Encode(),
	}

	app.render(w, r, "dog", blockData, http.StatusOK)
}

func (app *Application) dogTable(w http.ResponseWriter, r *http.Request) {
	colour := r.URL.Query().Get("colour")
	breed := r.URL.Query().Get("breed")

	if colour == "" {
		colour = "all"
	}

	if breed == "" {
		breed = "all"
	}

	ctx := r.Context()
	colours, err := app.dogRepo.Colours(ctx)
	if err != nil {
		app.serverError(w, r, fmt.Errorf("app.dogRepo.Colours(): %w", err), http.StatusInternalServerError)

		return
	}

	breeds, err := app.dogRepo.Breeds(ctx)
	if err != nil {
		app.serverError(w, r, fmt.Errorf("app.dogRepo.Colours(): %w", err), http.StatusInternalServerError)

		return
	}

	colours = append([]string{"all"}, colours...)
	breeds = append([]string{"all"}, breeds...)

	dogs, _ := app.dogRepo.All(ctx, colour, breed)

	blockData := map[string]any{
		"Colours": colours,
		"Breeds":  breeds,
		"Dogs":    dogs,

		"Colour": colour,
		"Breed":  breed,
	}

	w.Header().Set("Hx-Push-Url", "/dog?"+r.URL.Query().Encode())
	app.render(w, r, "dog-table", blockData, http.StatusOK)
}
