package checkgrp

import (
	"net/http"

	"github.com/youngjun827/thoughts/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	hdl := New(cfg.Build)
	app.HandleNoMiddleware(http.MethodGet, version, "/readiness", hdl.Readiness)
	app.HandleNoMiddleware(http.MethodGet, version, "/liveness", hdl.Liveness)
}