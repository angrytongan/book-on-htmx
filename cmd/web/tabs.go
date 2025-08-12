package main

import (
	"net/http"
)

type TabLink struct {
	Label  string
	Href   string
	Active bool
}

func (app *Application) tabs(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "tabs", nil, http.StatusOK)
}

func (app *Application) tabsRadioButtons(w http.ResponseWriter, r *http.Request) {
	tabs := []TabLink{
		{"One", "/tabs/radio-buttons/one", true},
		{"Two", "/tabs/radio-buttons/two", false},
		{"Three", "/tabs/radio-buttons/three", false},
	}

	pageData := map[string]any{
		"Tabs": tabs,
	}

	app.render(w, r, "tab-radio-buttons", pageData, http.StatusOK)
}

func (app *Application) tabsRadioButtonsTab(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("tab")
	block := "tab-content-" + tab

	app.render(w, r, block, nil, http.StatusOK)
}
