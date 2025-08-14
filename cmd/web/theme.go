package main

import (
	"context"
	"fmt"
	"net/http"
)

func (app *Application) theme(w http.ResponseWriter, r *http.Request) {
	id := 1

	activeTheme, err := app.themeRepo.Active(context.Background(), id)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("app.themeRepo.Active(%d): %w", id, err),
			http.StatusInternalServerError,
			"Couldn't load active theme!",
		)

		return
	}

	themes, err := app.themeRepo.Themes(r.Context(), activeTheme)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("app.themeRepo.Themes(%s): %w", activeTheme, err),
			http.StatusInternalServerError,
			"Couldn't load themes!",
		)
	}

	pageData := map[string]any{
		"Themes": themes,
	}

	app.render(w, r, "theme", pageData, http.StatusOK)
}

func (app *Application) themeChooserSave(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.serverError(w, r, fmt.Errorf("r.ParseForm(): %w", err), http.StatusInternalServerError)

		return
	}

	id := 1
	activeTheme := r.FormValue("theme")

	if err := app.themeRepo.SetActive(context.Background(), id, activeTheme); err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("theme.SetActive(%d, %s): %w", id, activeTheme, err),
			http.StatusInternalServerError,
		)

		return
	}

	app.clientRedirect(w, r, "/theme", http.StatusSeeOther)
}
