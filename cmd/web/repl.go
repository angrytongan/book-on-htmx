package main

import "net/http"

func (app *Application) repl(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "repl", nil, http.StatusOK)
}
