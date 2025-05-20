package main

import (
	"net/http"

	"github.com/gfffrtt/go-next/pkg/server"
	"github.com/gfffrtt/go-next/src/app"
	"github.com/gfffrtt/go-next/src/app/about"
	"github.com/gfffrtt/go-next/src/app/counter"
	"github.com/gfffrtt/go-next/src/app/search"
)

func main() {
	app := server.NewRouter().
		WithStatic().
		WithLayout(app.Layout).
		Page("/", app.Page).
		Page("/about", about.Page).
		Page("/search", search.Page).
		Page("/counter", counter.Page)

	http.ListenAndServe(":8080", app.App)
}
