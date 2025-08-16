package main

import (
	"context"
	"fmt"
	"net/http"
)

const (
	ErrTextCouldntLoadActiveTheme = "Couldn't load active theme!"
	ErrTextCouldntLoadThemes      = "Couldn't load themes!"
)

func (app *Application) theme(w http.ResponseWriter, r *http.Request) {
	id := 1

	activeTheme, errActiveTheme := app.themeRepo.Active(context.Background(), id)
	themes, errThemes := app.themeRepo.Themes(r.Context(), activeTheme)

	pageData := map[string]any{
		"Themes": themes,
	}

	if errActiveTheme != nil {
		pageData["Error"] = ErrTextCouldntLoadActiveTheme
	}

	if errThemes != nil {
		pageData["Error"] = ErrTextCouldntLoadThemes
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
