package main

import "net/http"

func (app *Application) root(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "root", nil, http.StatusOK)
}
