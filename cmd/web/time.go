package main

import (
	"net/http"
	"time"
)

func (app *Application) time(w http.ResponseWriter, r *http.Request) {
	blockData := map[string]any{
		"Time": time.Now().Format(time.DateTime),
	}

	app.render(w, r, "time", blockData, http.StatusOK)
}

func (app *Application) servertime(w http.ResponseWriter, r *http.Request) {
	blockData := map[string]any{
		"Time": time.Now().Format(time.DateTime),
	}

	app.render(w, r, "time-servertime", blockData, http.StatusOK)
}
