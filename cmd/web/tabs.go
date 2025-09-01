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

	blockData := map[string]any{
		"Tabs":    tabs,
		"Content": "This is content " + tab + ".",
	}

	app.render(w, r, "tab-links", blockData, http.StatusOK)
}

func (app *Application) tabsLinksOOB(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("tab")
	if tab == "" {
		tab = "one"
	}

	tabs := []LinkTabLink{
		{"One", "/tabs/links-oob/one", tab == "one"},
		{"Two", "/tabs/links-oob/two", tab == "two"},
		{"Three", "/tabs/links-oob/three", tab == "three"},
	}

	blockData := map[string]any{
		"Tabs":   tabs,
		"Active": tab,
	}

	app.render(w, r, "tab-links-oob", blockData, http.StatusOK)
}

func (app *Application) tabsButtonsOOB(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("tab")
	if tab == "" {
		tab = "one"
	}

	tabs := []ButtonTabLink{
		{"One", "/tabs/buttons-oob/one", tab == "one"},
		{"Two", "/tabs/buttons-oob/two", tab == "two"},
		{"Three", "/tabs/buttons-oob/three", tab == "three"},
	}

	blockData := map[string]any{
		"Tabs":   tabs,
		"Active": tab,
	}

	app.render(w, r, "tab-buttons-oob", blockData, http.StatusOK)
}

func (app *Application) tabsRadiosOOB(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("tab")
	if tab == "" {
		tab = "one"
	}

	tabs := []RadioTabLink{
		{"One", "one", tab == "one"},
		{"Two", "two", tab == "two"},
		{"Three", "three", tab == "three"},
	}

	blockData := map[string]any{
		"Tabs": tabs,
	}

	app.render(w, r, "tab-radio-oob", blockData, http.StatusOK)
}

func (app *Application) tabsContent(w http.ResponseWriter, r *http.Request) {
	tab := r.PathValue("content")
	block := "tab-content-" + tab

	app.render(w, r, block, nil, http.StatusOK)
}
