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
		// Configures the persistant layout for all pages
		WithLayout(app.Layout).
		Page("/", app.Page).
		Page("/about", about.Page).
		// This route contains PPR (Partial Prerendering)
		Page("/search", search.Page).
		// This route contains a client component
		Page("/counter", counter.Page)

	http.ListenAndServe(":8080", app.App)
}
