package server

import (
	"net/http"
	"os"

	"github.com/gfffrtt/go-next/pkg/html"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	App    *chi.Mux
	Layout func(children ...html.Element) html.Element
}

func (router *Router) Page(path string, handler func(r *http.Request) html.Element) *Router {
	router.App.Get(path, func(w http.ResponseWriter, r *http.Request) {
		stream := html.NewStream()

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")

		flusher := w.(http.Flusher)

		page := handler(r)

		if router.Layout != nil {
			page = router.Layout(page)
		}

		if path == "/" {
			path = "/index"
		}

		if _, err := os.Stat("./build/router/" + path); err == nil {
			page.AddChild(html.Script(map[string]string{
				"src":  "/static/router" + path + "/index.js",
				"type": "module",
			}))
		}

		w.Write([]byte(page.Render(stream)))

		flusher.Flush()

		go func() {
			for html := range stream.Channel {
				w.Write([]byte(html))
				flusher.Flush()
			}
		}()

		stream.Wait()
	})
	return router
}

func NewRouter() *Router {
	return &Router{
		App: chi.NewMux(),
	}
}

func (r *Router) WithLayout(layout func(children ...html.Element) html.Element) *Router {
	r.Layout = layout
	return &Router{
		App:    r.App,
		Layout: layout,
	}
}

func (r *Router) WithStatic() *Router {
	r.App.Handle(
		"/static/*",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./build"))),
	)
	return r
}
