package main

import (
	"fmt"
	"net/http"
)

func (app *Application) themeChooser(w http.ResponseWriter, r *http.Request) {
	id := 1

	activeTheme, err := app.themeRepo.Active(id)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("app.themeRepo.Active(%d): %w", id, err),
			http.StatusInternalServerError,
		)
	}

	themes, _ := app.themeRepo.Themes(activeTheme)

	pageData := map[string]any{
		"Themes": themes,
	}

	app.render(w, r, "theme-chooser", pageData, http.StatusOK)
}

func (app *Application) themeChooserSave(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.serverError(w, r, fmt.Errorf("r.ParseForm(): %w", err), http.StatusInternalServerError)

		return
	}

	id := 1
	activeTheme := r.FormValue("theme")

	if err := app.themeRepo.SetActive(id, activeTheme); err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("theme.SetActive(%d, %s): %w", id, activeTheme, err),
			http.StatusInternalServerError,
		)

		return
	}

	app.clientRedirect(w, r, "/settings", http.StatusSeeOther)
}
