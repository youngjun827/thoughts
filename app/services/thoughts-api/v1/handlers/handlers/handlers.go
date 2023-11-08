// Package handlers add all route groups.
package handlers

import (
	"github.com/youngjun827/thoughts/app/services/thoughts-api/v1/handlers/hackgrp"
	v1 "github.com/youngjun827/thoughts/business/web/v1"
	"github.com/youngjun827/thoughts/foundation/web"
)

type Routes struct{}

func (r Routes) Add(app *web.App, cfg v1.APIMuxConfig) {
	hackgrp.Routes(app)
}
