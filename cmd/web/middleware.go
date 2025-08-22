package main

import (
	"fmt"
	"net/http"
	"time"
)

func delayResponse(ms int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Duration(ms) * time.Millisecond)

			next.ServeHTTP(w, r)
		})
	}
}

func (app *Application) mustBeHtmxRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Hx-Request") != "true" {
			app.serverError(w, r, fmt.Errorf("not a hx-request"), http.StatusBadRequest)

			return
		}

		next.ServeHTTP(w, r)
	})
}
