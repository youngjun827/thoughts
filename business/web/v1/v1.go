// Package v1 manages the different versions of the API.
package v1

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/youngjun827/thoughts/business/web/v1/mid"
	"github.com/youngjun827/thoughts/foundation/logger"
	"github.com/youngjun827/thoughts/foundation/web"
)

type APIMuxConfig struct {
    Build    string
    Shutdown chan os.Signal
    Log      *logger.Logger
}

type RouteAdder interface {
    Add(router chi.Router, cfg APIMuxConfig)
}

func APIMux(cfg APIMuxConfig, routeAdder RouteAdder) chi.Router {
    app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

    routeAdder.Add(app.Mux, cfg)

    return app.Mux
}
