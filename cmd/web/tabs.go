package main

import (
	"net/http"
)

type LinkTabLink struct {
	Label  string
	Href   string
	Active bool
}

type ButtonTabLink struct {
	Label  string
	Href   string
	Active bool
}

type RadioTabLink struct {
	Label   string
	Content string
	Active  bool
}

func (app *Application) tabs(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "tabs", nil, http.StatusOK)
}

func (app *Application) tabsLinks(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("tab")
	if tab == "" {
		tab = "one"
	}

	tabs := []LinkTabLink{
		{"One", "/tabs/links/one", tab == "one"},
		{"Two", "/tabs/links/two", tab == "two"},
		{"Three", "/tabs/links/three", tab == "three"},
	}

	pageData := map[string]any{
		"Tabs":   tabs,
		"Active": tab,
	}

	app.render(w, r, "tab-links", pageData, http.StatusOK)
}

func (app *Application) tabsButtons(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("tab")
	if tab == "" {
		tab = "one"
	}

	tabs := []ButtonTabLink{
		{"One", "/tabs/buttons/one", tab == "one"},
		{"Two", "/tabs/buttons/two", tab == "two"},
		{"Three", "/tabs/buttons/three", tab == "three"},
	}

	pageData := map[string]any{
		"Tabs":   tabs,
		"Active": tab,
	}

	app.render(w, r, "tab-buttons", pageData, http.StatusOK)
}

func (app *Application) tabsRadioButtons(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("tab")
	if tab == "" {
		tab = "one"
	}

	tabs := []RadioTabLink{
		{"One", "one", tab == "one"},
		{"Two", "two", tab == "two"},
		{"Three", "three", tab == "three"},
	}

	pageData := map[string]any{
		"Tabs": tabs,
	}

	app.render(w, r, "tab-radio-buttons", pageData, http.StatusOK)
}

func (app *Application) tabsContent(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("content")
	block := "tab-content-" + tab

	app.render(w, r, block, nil, http.StatusOK)
}
