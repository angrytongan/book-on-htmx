package main

import "net/http"

func (app *Application) form(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "form", nil, http.StatusOK)
}

func (app *Application) formProcess(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
