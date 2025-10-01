package main

import (
	"errors"
	"net/http"
	"slices"
)

func (app *Application) settings(w http.ResponseWriter, r *http.Request) {
	tab := r.URL.Query().Get("tab")
	if tab == "" {
		app.clientLocation(w, r, "/settings?tab=first&update-nav=true", http.StatusSeeOther)

		return
	}

	if !slices.Contains([]string{"first", "second", "third"}, tab) {
		app.serverError(w, r, errors.New("invalid tab"), http.StatusBadRequest)

		return
	}

	tabs := []LinkTabLink{
		{"First", "/settings?tab=first", tab == "first"},
		{"Second", "/settings?tab=second", tab == "second"},
		{"Third", "/settings?tab=third", tab == "third"},
	}

	var content map[string]any

	switch tab {
	case "first":
		content = settingsFirst()

	case "second":
		content = settingsSecond()

	case "third":
		content = settingsThird()
	}

	blockData := map[string]any{
		"Tab":     tab,
		"Tabs":    tabs,
		"Content": content,
	}

	var activeTab LinkTabLink
	for _, activeTab = range tabs {
		if activeTab.Active {
			break
		}
	}

	w.Header().Add("Hx-Push-Url", activeTab.Href)

	if r.URL.Query().Get("update-nav") != "" {
		app.render(w, r, "settings", blockData, http.StatusOK)
	} else {
		app.render(w, r, "settings-tab", blockData, http.StatusOK)
	}
}

func settingsFirst() map[string]any {
	return map[string]any{
		"Content": "this is settingsFirst()",
	}
}

func settingsSecond() map[string]any {
	return map[string]any{
		"Content": "this is settingsSecond()",
	}
}

func settingsThird() map[string]any {
	return map[string]any{
		"Content": "this is settingsThird()",
	}
}
