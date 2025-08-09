package main

const (
	port = 8880
)

func main() {
	app := newApplication()

	// Navigation.
	app.mux.Get("/nav/", app.nav)
	app.mux.Get("/nav/{href}", app.nav)

	// Pages.
	app.mux.Get("/", app.home)
	app.mux.Get("/tabs", app.tabs)
	app.mux.Get("/settings", app.settings)

	_ = app.listen(port)
}
