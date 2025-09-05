package main

import "net/http"

func (app *Application) drawer(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "drawer", nil, http.StatusOK)
}
