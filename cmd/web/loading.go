package main

import (
	"math/rand"
	"net/http"
	"time"
)

const (
	minDelay = 5000  // ms
	maxDelay = 10000 // ms
)

func (app *Application) loading(w http.ResponseWriter, r *http.Request) {
	blockData := map[string]any{
		"NumLoaders": 54,
	}

	app.render(w, r, "loading", blockData, http.StatusOK)
}

func (app *Application) loadingThing(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(time.Second)
	number := r.PathValue("number")

	delay := rand.Intn(maxDelay-minDelay+1) + minDelay

	blockData := map[string]any{
		"Number": number,
		"Delay":  delay,
		"Value":  time.Now().Format(time.DateTime + " MST"),
	}

	app.render(w, r, "loading-thing", blockData, http.StatusOK)
}
