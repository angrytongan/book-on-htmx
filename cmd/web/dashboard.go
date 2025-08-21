package main

import "net/http"

func (app *Application) dashboard(w http.ResponseWriter, r *http.Request) {
	widgets := []string{"/weather", "/time"}

	pageData := map[string]any{
		"Widgets": widgets,
	}

	app.renderWithNav(w, r, "dashboard", pageData, http.StatusOK)
}
