package main

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
)

const (
	maxRows = 100
)

var (
	ErrInvalidColour = errors.New("invalid colour")
)

type Dog struct {
	Colour string
	Breed  string
	Name   string
}

func (app *Application) dog(w http.ResponseWriter, r *http.Request) {
	blockData := map[string]any{
		"QP": r.URL.Query().Encode(),
	}

	app.render(w, r, "dog", blockData, http.StatusOK)
}

func (app *Application) dogTable(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	colours, err := app.dogRepo.Colours(ctx)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("app.dogRepo.Colours(): %w", err),
			http.StatusInternalServerError,
		)

		return
	}

	breeds, err := app.dogRepo.Breeds(ctx)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf("app.dogRepo.Colours(): %w", err),
			http.StatusInternalServerError,
		)

		return
	}

	colour := r.URL.Query().Get("colour")
	breed := r.URL.Query().Get("breed")
	orderDirection := r.URL.Query().Get("order-direction")
	lastOrderDirection := r.URL.Query().Get("last-order-direction")

	if colour == "" {
		colour = "all"
	} else if colour != "all" && !slices.Contains(colours, colour) {
		app.serverError(
			w,
			r,
			fmt.Errorf("%s: %w", colour, ErrInvalidColour),
			http.StatusBadRequest,
			"Invalid colour",
		)

		return
	}

	if breed == "" {
		breed = "all"
	} else if breed != "all" && !slices.Contains(breeds, breed) {
		app.serverError(w, r, fmt.Errorf("invalid breed %s", colour), http.StatusBadRequest, "Invalid breed")

		return
	}

	order := ""
	direction := ""
	queryOrderDirection := "colour ASC, breed ASC, name ASC"

	if orderDirection != "" {
		// Radio button was used to determine order and column. Use orderDirection
		// to figure out how to sort our data.
		tokens := strings.Split(orderDirection, "-")
		order = tokens[0]
		direction = tokens[1]
		queryOrderDirection = order + " " + direction

		// Remember this order and direction on the page in case the next hit
		// doesn't come from the radio button.
		lastOrderDirection = orderDirection
	} else if lastOrderDirection != "" {
		// Radio button wasn't used to order - some other control asked for the
		// refresh. Use lastOrderDirection to give us the order and direction that
		// was used last time. Retain it for render.
		tokens := strings.Split(lastOrderDirection, "-")
		order = tokens[0]
		direction = tokens[1]
		queryOrderDirection = order + " " + direction
	}

	limit := maxRows

	colours = append([]string{"all"}, colours...)
	breeds = append([]string{"all"}, breeds...)

	dogs, err := app.dogRepo.All(ctx, colour, breed, queryOrderDirection, limit)
	if err != nil {
		app.serverError(
			w,
			r,
			fmt.Errorf(
				"app.dogRepo.All(%s, %s, %s, %d): %w",
				colour,
				breed,
				queryOrderDirection,
				limit,
				err,
			),
			http.StatusInternalServerError,
		)

		return
	}

	blockData := map[string]any{
		"Colours": colours,
		"Breeds":  breeds,
		"Dogs":    dogs,
		"Colour":  colour,
		"Breed":   breed,

		"Order":              order,
		"Direction":          direction,
		"LastOrderDirection": lastOrderDirection,
	}

	w.Header().Set("Hx-Push-Url", "/dog?"+r.URL.Query().Encode())
	app.render(w, r, "dog-table", blockData, http.StatusOK)
}
