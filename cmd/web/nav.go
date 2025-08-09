package main

import (
	"net/http"

	"bonh/internal/nav"
)

func (app *Application) nav(w http.ResponseWriter, r *http.Request) {
	path := "/" + r.PathValue("href")

	pageData := map[string]any{
		"Nav":  nav.PageLinks(path),
		"Href": path,
	}

	w.Header().Add("Hx-Push-Url", path)
	app.render(w, r, "nav-main", pageData, http.StatusOK)
}
