package main

import "net/http"

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	app.renderWithNav(w, r, "home", nil, http.StatusOK)
}
