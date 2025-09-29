package main

import "net/http"

func (app *Application) mapForm(w http.ResponseWriter, r *http.Request) {
	blockData := map[string]any{}

	blockData["CC"] = "au"
	blockData["NextDisabled"] = true

	app.render(w, r, "map-form", blockData, http.StatusOK)
}

func (app *Application) mapFormForm(w http.ResponseWriter, r *http.Request) {
	blockData := map[string]any{}

	blockData["CC"] = "au"
	blockData["NextDisabled"] = true

	app.render(w, r, "map-form-form", blockData, http.StatusOK)
}
