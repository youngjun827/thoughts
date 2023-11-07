package v1

import (
	"os"

	"github.com/go-chi/chi/v5"
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
    app := web.NewApp(cfg.Shutdown)

    routeAdder.Add(app.Mux, cfg)

    return app.Mux
}
