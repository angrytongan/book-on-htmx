package main

import "net/http"

func (app *Application) settings(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "settings", nil, http.StatusOK)
}
