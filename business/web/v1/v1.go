// Package v1 manages the different versions of the API.
package v1

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/youngjun827/thoughts/business/web/v1/mid"
	"github.com/youngjun827/thoughts/foundation/logger"
	"github.com/youngjun827/thoughts/foundation/web"
)

type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
	DB 		 *sqlx.DB
}
type RouteAdder interface {
	Add(app *web.App, cfg APIMuxConfig)
}

func APIMux(cfg APIMuxConfig, routeAdder RouteAdder) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())

	routeAdder.Add(app, cfg)

	return app
}
