package main

import "net/http"

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home", nil, http.StatusOK)
}
