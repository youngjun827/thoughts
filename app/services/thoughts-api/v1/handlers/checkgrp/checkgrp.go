// Package checkgrp adds specific routes for checking app status in k8s.
package checkgrp

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	db "github.com/youngjun827/thoughts/business/database/dbsql/pgx"
	"github.com/youngjun827/thoughts/foundation/logger"
	"github.com/youngjun827/thoughts/foundation/web"
)

type Handlers struct {
	build string
	log   *logger.Logger
	db    *sqlx.DB
}

func New(build string, log *logger.Logger, db *sqlx.DB) *Handlers {
	return &Handlers{
		build: build,
		log:   log,
		db:    db,
	}
}

func (h *Handlers) Readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	status := "OK"
	statusCode := http.StatusOK
	err := db.StatusCheck(ctx, h.db)
	if err != nil {
		status = "db not ready"
		statusCode = http.StatusInternalServerError
		h.log.Info(ctx, "readiness failure", "status", status)
	}

	data := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	h.log.Info(ctx, "readiness", "status", status)

	return web.Respond(ctx, w, data, statusCode)
}

func (h *Handlers) Liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	data := struct {
		Status     string `json:"status,omitempty"`
		Build      string `json:"build,omitempty"`
		Host       string `json:"host,omitempty"`
		Name       string `json:"name,omitempty"`
		PodIP      string `json:"podIP,omitempty"`
		Node       string `json:"node,omitempty"`
		Namespace  string `json:"namespace,omitempty"`
		GOMAXPROCS string `json:"GOMAXPROCS,omitempty"`
	}{
		Status:     "up",
		Build:      h.build,
		Host:       host,
		Name:       os.Getenv("KUBERNETES_NAME"),
		PodIP:      os.Getenv("KUBERNETES_POD_IP"),
		Node:       os.Getenv("KUBERNETES_NODE_NAME"),
		Namespace:  os.Getenv("KUBERNETES_NAMESPACE"),
		GOMAXPROCS: os.Getenv("GOMAXPROCS"),
	}

	h.log.Info(ctx, "liveness", "status", "OK")

	return web.Respond(ctx, w, data, http.StatusOK)
}
