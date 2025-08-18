package main

import (
	"fmt"
	"net/http"
)

const (
	ErrTextFailedSearch = "Failed to search database!"
)

func (app *Application) search(w http.ResponseWriter, r *http.Request) {
	app.renderWithNav(w, r, "search", nil, http.StatusOK)
}

func (app *Application) searchTerm(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	block := "search-results"
	results := []string{}

	if term == "" {
		block = "search-instructions"
	} else {
		var err error

		results, err = app.searchRepo.Term(r.Context(), term)

		if err != nil {
			app.serverError(
				w,
				r,
				fmt.Errorf("app.searchRepo.Term(%s): %w", term, err),
				http.StatusInternalServerError,
				ErrTextFailedSearch,
			)

			return
		}

		if len(results) == 0 {
			block = "search-no-results"
		}
	}

	pageData := map[string]any{
		"Results": results,
	}

	app.renderPartial(w, r, block, pageData, http.StatusOK)
}
