package main

import (
	"math/rand"
	"net/http"
	"time"
)

func (app *Application) toast(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "toast", nil, http.StatusOK)
}

func (app *Application) toastServerTime(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format(time.DateTime)

	blockData := map[string]any{
		"Now": now,
	}

	app.render(w, r, "toast-server-time", blockData, http.StatusOK)
}

func (app *Application) toastRandomNumber(w http.ResponseWriter, r *http.Request) {
	randomNum := rand.Intn(101)

	blockData := map[string]any{
		"Number": randomNum,
	}

	app.render(w, r, "toast-random-number", blockData, http.StatusOK)
}

func (app *Application) toastRandomLetter(w http.ResponseWriter, r *http.Request) {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	randomLetter := string(letters[rand.Intn(len(letters))])

	blockData := map[string]any{
		"Letter": randomLetter,
	}

	app.render(w, r, "toast-random-letter", blockData, http.StatusOK)
}

func (app *Application) toastRandomWord(w http.ResponseWriter, r *http.Request) {
	if len(app.words) == 0 {
		blockData := map[string]any{
			"Message": http.StatusText(http.StatusInternalServerError),
		}
		app.render(w, r, "toast-error", blockData, http.StatusInternalServerError)

		return
	}

	randomWord := app.words[rand.Intn(len(app.words))]

	blockData := map[string]any{
		"Word": randomWord,
	}

	app.render(w, r, "toast-random-word", blockData, http.StatusOK)
}
