// Package hackgrp adds specific routes for hackgrp.
package hackgrp

import (
	"net/http"

	"github.com/youngjun827/thoughts/foundation/web"
)

func Routes(app *web.App) {
	app.Handle(http.MethodGet, "/hack", Hack)
}
