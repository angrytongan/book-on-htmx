package main

import (
	"bonh/internal/search"
	"net/http"
)

func (app *Application) search(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "search", nil, http.StatusOK)
}

func (app *Application) searchTerm(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("term")
	block := "search-results"
	results := []string{}

	if term == "" {
		block = "search-instructions"
	} else {
		results = search.Term(term)

		if len(results) == 0 {
			block = "search-no-results"
		}
	}

	pageData := map[string]any{
		"Results": results,
	}

	app.render(w, r, block, pageData, http.StatusOK)
}
