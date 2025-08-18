package main

import (
	"net/http"
	"time"
)

func (app *Application) toast(w http.ResponseWriter, r *http.Request) {
	app.renderPage(w, r, "toast", nil, http.StatusOK)
}

func (app *Application) toastServerTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.DateTime)

	pageData := map[string]any{
		"Now": now,
	}

	app.renderPartial(w, r, "toast-server-time", pageData, http.StatusOK)
}
