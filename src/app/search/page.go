package search

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gfffrtt/go-next/pkg/html"
)

type data struct {
	Name string
	Age  int
	City string
}

var names = []data{
	{Name: "John", Age: 20, City: "New York"},
	{Name: "Jane", Age: 21, City: "Los Angeles"},
	{Name: "Jim", Age: 22, City: "Chicago"},
}

func searchData(search string) []data {
	results := []data{}

	for _, d := range names {
		if strings.Contains(strings.ToLower(strings.TrimSpace(d.Name)), strings.ToLower(strings.TrimSpace(search))) {
			results = append(results, d)
		}
	}

	time.Sleep(5 * time.Second)

	return results
}

func Page(r *http.Request) html.Element {
	q := r.URL.Query().Get("q")
	return html.Fragment(
		html.P(map[string]string{}, html.String(fmt.Sprintf("Search results for %s", q))),
		html.Suspense(func() html.Element {
			results := searchData(q)
			return html.Ul(
				map[string]string{},
				html.Map(results, func(result data) html.Element {
					return html.Li(map[string]string{}, html.String(result.Name))
				})...,
			)
		}, html.P(map[string]string{}, html.String("Loading..."))),
	)
}
