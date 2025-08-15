package main

import (
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
