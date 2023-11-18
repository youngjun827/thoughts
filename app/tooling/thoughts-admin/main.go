package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ardanlabs/conf/v3"
	"github.com/youngjun827/thoughts/business/data/dbmigrate"
	db "github.com/youngjun827/thoughts/business/data/dbsql/pgx"
)


func main() {
	err := migrateSeed()

	if err != nil {
		log.Fatalln(err)
	}
}

func migrateSeed() error {
	var cfg struct {
		DB struct {
			User         string `conf:"default:postgres"`
			Password     string `conf:"default:postgres,mask"`
			Host         string `conf:"default:database-thoughts.thoughts-system.svc.cluster.local"`
			Name         string `conf:"default:postgres"`
			MaxIdleConns int    `conf:"default:2"`
			MaxOpenConns int    `conf:"default:0"`
			DisableTLS   bool   `conf:"default:true"`
		}
	}

	const prefix = "thoughts"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}

		return fmt.Errorf("parsing config: %w", err)
	}

	dbConfig := db.Config{
		User:         cfg.DB.User,
		Password:     cfg.DB.Password,
		Host:         cfg.DB.Host,
		Name:         cfg.DB.Name,
		MaxIdleConns: cfg.DB.MaxIdleConns,
		MaxOpenConns: cfg.DB.MaxOpenConns,
		DisableTLS:   cfg.DB.DisableTLS,
	}

	db, err := db.Open(dbConfig)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = dbmigrate.Migrate(ctx, db)
	if err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")

	// -------------------------------------------------------------------------

	err = dbmigrate.Seed(ctx, db)
	if err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")
	return nil
}