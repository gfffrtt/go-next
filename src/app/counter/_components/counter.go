package _components

import "github.com/gfffrtt/go-next/pkg/html"

func Counter() html.Element {
	return html.Client(
		"counter",
		"/counter/_components/counter.tsx",
		map[string]int{"count": 5},
	)
}
