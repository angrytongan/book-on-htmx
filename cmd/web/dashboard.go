package main

import "net/http"

func (app *Application) dashboard(w http.ResponseWriter, r *http.Request) {
	widgets := []string{"/weather", "/time"}

	blockData := map[string]any{
		"Widgets": widgets,
	}

	app.render(w, r, "dashboard", blockData, http.StatusOK, "navigation")
}
