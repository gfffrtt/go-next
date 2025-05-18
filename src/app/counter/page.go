package counter

import (
	"net/http"

	"github.com/gfffrtt/go-next/pkg/html"
	"github.com/gfffrtt/go-next/src/app/counter/_components"
)

func Page(r *http.Request) html.Element {
	return html.Fragment(
		html.Client(
			"counter",
			"/counter/_components/counter.tsx",
			map[string]int{"count": 5},
		),
		_components.Counter(),
	)
}
