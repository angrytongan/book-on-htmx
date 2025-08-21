package main

import (
	"net/http"
	"time"
)

func (app *Application) weather(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "weather", nil, http.StatusOK)
}

func (app *Application) weatherCurrent(w http.ResponseWriter, r *http.Request) {
	pageData := map[string]any{
		"Temperature":   16,
		"WindDirection": "NE",
		"WindSpeed":     20,
		"Conditions":    "Sunny",
		"ServerTime":    time.Now().Format(time.DateTime),
	}

	app.render(w, r, "weather-data", pageData, http.StatusOK)
}
