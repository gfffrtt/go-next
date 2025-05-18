package app

import (
	"net/http"

	"github.com/gfffrtt/go-next/pkg/html"
)

func Page(r *http.Request) html.Element {
	return html.Form(
		map[string]string{
			"class":  "flex gap-5",
			"action": "/search",
			"method": "GET",
		},
		html.Input(map[string]string{"type": "text", "placeholder": "Search", "name": "q", "id": "q", "class": "border rounded-md px-2 py-1"}),
		html.Button(map[string]string{"type": "submit", "class": "border rounded-md px-2 py-1 bg-black text-white"}, html.String("Search")),
	)
}
