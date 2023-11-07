// Add all route groups.
package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/youngjun827/thoughts/app/services/thoughts-api/v1/handlers/hackgrp"
	v1 "github.com/youngjun827/thoughts/business/web/v1"
)


type Routes struct{}

func (r Routes) Add(router chi.Router, cfg v1.APIMuxConfig) {
    hackgrp.Routes(router)
}