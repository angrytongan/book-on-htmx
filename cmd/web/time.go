package main

import (
	"net/http"
	"time"
)

func (app *Application) time(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Time": time.Now().Format(time.DateTime),
	}

	app.render(w, r, "time", pageData, http.StatusOK)
}

func (app *Application) servertime(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Time": time.Now().Format(time.DateTime),
	}

	app.render(w, r, "time-servertime", pageData, http.StatusOK)
}
