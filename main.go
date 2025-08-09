package main

const (
	port = 8880
)

func main() {
	app := newApplication()

	app.mux.Get("/", app.home)
	app.mux.Get("/tabs", app.tabs)
	app.mux.Get("/settings", app.settings)

	app.listen(port)
}
