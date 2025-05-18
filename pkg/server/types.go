package server

import (
	"net/http"

	"github.com/gfffrtt/go-next/pkg/html"
)

type Page func(r *http.Request) html.Element
