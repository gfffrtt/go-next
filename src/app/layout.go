package app

import "github.com/gfffrtt/go-next/pkg/html"

func Layout(children ...html.Element) html.Element {
	return html.Html(
		map[string]string{"lang": "en"},
		html.Head(
			map[string]string{},
			html.Meta(map[string]string{"charset": "utf-8"}),
			html.Meta(map[string]string{"name": "viewport", "content": "width=device-width, initial-scale=1.0"}),
			html.Title(map[string]string{}, html.String("Go Next")),
			html.Link(map[string]string{"rel": "stylesheet", "href": "https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css"}),
		),
		html.Body(
			map[string]string{"class": "flex flex-col gap-5"},
			html.Header(
				map[string]string{},
				html.Nav(
					map[string]string{"class": "w-full flex justify-start items-center gap-5"},
					html.A(map[string]string{"class": "border px-2 py-1 rounded-md border font-bold", "href": "/"}, html.String("Home")),
					html.A(map[string]string{"class": "border px-2 py-1 rounded-md border font-bold", "href": "/about"}, html.String("About")),
					html.A(map[string]string{"class": "border px-2 py-1 rounded-md border font-bold", "href": "/counter"}, html.String("Counter")),
				),
			),
			html.Main(
				map[string]string{},
				children...,
			),
		),
	)
}
