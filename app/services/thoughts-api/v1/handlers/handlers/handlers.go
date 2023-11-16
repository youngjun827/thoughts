// Package handlers add all route groups.
package handlers

import (
	"github.com/youngjun827/thoughts/app/services/thoughts-api/v1/handlers/bloggrp"
	"github.com/youngjun827/thoughts/app/services/thoughts-api/v1/handlers/checkgrp"
	"github.com/youngjun827/thoughts/app/services/thoughts-api/v1/handlers/hackgrp"
	v1 "github.com/youngjun827/thoughts/business/web/v1"
	"github.com/youngjun827/thoughts/foundation/web"
)

type Routes struct{}

func (Routes) Add(app *web.App, apiCfg v1.APIMuxConfig) {
	hackgrp.Routes(app)

	checkgrp.Routes(app, checkgrp.Config{
		Build: apiCfg.Build,
		Log: apiCfg.Log,
	})


	bloggrp.Routes(app, bloggrp.Config{
		Build: apiCfg.Build,
		Log:   apiCfg.Log,
		DB:    apiCfg.DB,
	})
}