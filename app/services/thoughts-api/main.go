package main

import (
	"context"
	"errors"
	"expvar"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/youngjun827/thoughts/app/services/thoughts-api/v1/handlers"
	db "github.com/youngjun827/thoughts/business/data/dbsql/pgx"
	v1 "github.com/youngjun827/thoughts/business/web/v1"
	"github.com/youngjun827/thoughts/business/web/v1/debug"
	"github.com/youngjun827/thoughts/foundation/logger"
	"github.com/youngjun827/thoughts/foundation/web"
)

var build = "develop"

func main() {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT ******")
		},
	}

	traceIDFunc := func(ctx context.Context) string {
		return web.GetTraceID(ctx)
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "THOUGHTS-API", traceIDFunc, events)

	// -------------------------------------------------------------------------

	ctx := context.Background()

	err := run(ctx, log)
	if err != nil {
		log.Error(ctx, "startup", "msg", err)
		return
	}
}

func run(ctx context.Context, log *logger.Logger) error {

	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0), "build", build)

	// -------------------------------------------------------------------------
	// Configuration

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration `conf:"default:5s"`
			WriteTimeout    time.Duration `conf:"default:10s"`
			IdleTimeout     time.Duration `conf:"default:60s"`
			ShutdownTimeout time.Duration `conf:"default:10s"`
			APIHost         string        `conf:"default:0.0.0.0:3000"`
			DebugHost       string        `conf:"default:0.0.0.0:4000"`
		}
		DB struct {
			User         string `conf:"default:postgres"`
			Password     string `conf:"default:postgres,mask"`
			Host         string `conf:"default:database-thoughts.thoughts-system.svc.cluster.local"`
			Name         string `conf:"default:postgres"`
			MaxIdleConns int    `conf:"default:2"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "YOUNG JUN JOO",
		},
	}

	const prefix = "THOUGHTS"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// -------------------------------------------------------------------------
	// App Starting

	log.Info(ctx, "starting service", "version", build)
	defer log.Info(ctx, "shutdown complete")

	out, err := conf.String(&cfg)
	if err != nil {
		return fmt.Errorf("generating config for output: %w", err)
	}
	log.Info(ctx, "startup", "config", out)

	expvar.NewString("build").Set(build)

	// -------------------------------------------------------------------------
	// Start Debug Service

	go func() {
		log.Info(ctx, "startup", "status", "debug v1 router started", "host", cfg.Web.DebugHost)

		err := http.ListenAndServe(cfg.Web.DebugHost, debug.Mux())
		if err != nil {
			log.Error(ctx, "shutdown", "status", "debug v1 router closed", "host", cfg.Web.DebugHost, "msg", err)
		}
	}()

	// -------------------------------------------------------------------------
	// Database Support

	log.Info(ctx, "startup", "status", "initializing database support", "host", cfg.DB.Host)

	db, err := db.Open(db.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	})
	if err != nil {
		return fmt.Errorf("connecting to db: %w", err)
	}
	defer func() {
		log.Info(ctx, "shutdown", "status", "stopping database support", "host", cfg.DB.Host)
		db.Close()
	}()

	// -------------------------------------------------------------------------
	// Start API Service

	log.Info(ctx, "startup", "status", "initializing V1 API support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	cfgMux := v1.APIMuxConfig{
		Build:    build,
		Shutdown: shutdown,
		Log:      log,
		DB:       db,
	}

	apiMux := v1.APIMux(cfgMux, handlers.Routes{})

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      apiMux,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     logger.NewStdLogger(log, logger.LevelError),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info(ctx, "startup", "status", "api router started", "host", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)
		defer log.Info(ctx, "shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(ctx, cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
