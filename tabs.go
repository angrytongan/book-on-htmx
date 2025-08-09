package main

import "net/http"

func (app *Application) tabs(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "tabs", nil, http.StatusOK)
}
