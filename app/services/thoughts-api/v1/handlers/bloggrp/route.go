package bloggrp

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/youngjun827/thoughts/business/core/blog"
	"github.com/youngjun827/thoughts/business/core/blog/stores/blogdb"
	"github.com/youngjun827/thoughts/foundation/logger"
	"github.com/youngjun827/thoughts/foundation/web"
)

type Config struct {
	Build string
	Log   *logger.Logger
	DB    *sqlx.DB
}

func Routes(app *web.App, cfg Config) {
	const version = "v1"

	usrCore := blog.NewCore(cfg.Log, blogdb.NewStore(cfg.Log, cfg.DB))

	hdl := New(usrCore)
	app.Handle(http.MethodPost, version, "/blogs", hdl.Create)
}