package main

import "net/http"

func (app *Application) leaflet(w http.ResponseWriter, r *http.Request) {
	app.renderPage(w, r, "leaflet", nil, http.StatusOK)
}
