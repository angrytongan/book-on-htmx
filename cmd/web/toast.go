package main

import (
	"math/rand"
	"net/http"
	"time"
)

func (app *Application) toast(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "toast", nil, http.StatusOK)
}

func (app *Application) toastServerTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.DateTime)

	blockData := map[string]any{
		"Now": now,
	}

	app.render(w, r, "toast-server-time", blockData, http.StatusOK)
}

func (app *Application) toastRandomNumber(w http.ResponseWriter, r *http.Request) {
	randomNum := rand.Intn(101)

	blockData := map[string]any{
		"Number": randomNum,
	}

	app.render(w, r, "toast-random-number", blockData, http.StatusOK)
}
