package main

const (
	port = 8880
)

func main() {
	app := newApplication()

	app.mux.Get("/", app.root)

	app.listen(port)
}
