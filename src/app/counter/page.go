package counter

import (
	"net/http"

	"github.com/gfffrtt/go-next/pkg/html"
)

func Page(r *http.Request) html.Element {
	return html.Fragment(
		html.Client(
			"counter",
			"/counter/_components/counter.tsx",
			map[string]int{"count": 5},
		),
	)
}
