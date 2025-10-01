package main

import (
	"net/http"
	"net/url"
)

func (app *Application) drawer(w http.ResponseWriter, r *http.Request) {
	blockData := map[string]any{
		"Content": []string{
			url.QueryEscape("this is the first content"),
			url.QueryEscape("this is the second lot of content"),
			url.QueryEscape("this is the third lot of content"),
		},
	}

	app.render(w, r, "drawer", blockData, http.StatusOK)
}

func (app *Application) drawerContent(w http.ResponseWriter, r *http.Request) {
	stuff := r.URL.Query().Get("q")
	blockData := map[string]any{
		"Stuff": stuff,
	}

	app.render(w, r, "drawer-content", blockData, http.StatusOK)
}
