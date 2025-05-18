package about

import (
	"net/http"

	"github.com/gfffrtt/go-next/pkg/html"
)

func Page(r *http.Request) html.Element {
	return html.P(map[string]string{}, html.String("About page"))
}
