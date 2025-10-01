package main

import (
	"fmt"
	"net/http"
)

func (app *Application) form(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "form", nil, http.StatusOK)
}

func (app *Application) formStep1(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "form-step1", nil, http.StatusOK)
}

func (app *Application) formStep1Process(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("r.ParseForm(): %w", err),
			http.StatusInternalServerError,
		)

		return
	}

	errors := map[string]any{}

	// "Validate"
	name := r.FormValue("name")
	password := r.FormValue("password")
	if name != "foo" {
		errors["Name"] = "invalid name (try 'foo')"
	}
	if password != "bar" {
		errors["Password"] = "invalid password (try 'bar')"
	}

	// If we have errors, swap in the error response template.
	if len(errors) > 0 {
		w.Header().Add("Hx-Reswap", "none")
		app.render(w, r, "form-step1-errors", errors, http.StatusOK)

		return
	}

	// No errors, carry on.
	err := app.formAdvance(w, r, "/form/step2", app.formStep2)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("app.formAdvance(/form/step2): %w", err),
			http.StatusInternalServerError,
		)

		return
	}
}

func (app *Application) formStep2(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "form-step2", nil, http.StatusOK)
}

func (app *Application) formAdvance(w http.ResponseWriter, r *http.Request, url string, fn http.HandlerFunc) error {
	req, err := http.NewRequestWithContext(r.Context(), "GET", "/form/step2", nil)
	if err != nil {
		return fmt.Errorf("http.NewRequestWithContext(): %w", err)
	}

	req.Header.Add("Hx-Request", "true")
	fn(w, req)

	return nil
}
