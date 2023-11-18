// Package dbmigrate provides support for database migrations and seeding.
package dbmigrate

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/ardanlabs/darwin/v3"
	"github.com/ardanlabs/darwin/v3/dialects/postgres"
	"github.com/ardanlabs/darwin/v3/drivers/generic"
	"github.com/jmoiron/sqlx"
	db "github.com/youngjun827/thoughts/business/database/dbsql/pgx"
)

var (
	//go:embed sql/migrate.sql
	migrateDoc string

	//go:embed sql/seed.sql
	seedDoc string
)

func Migrate(ctx context.Context, sqldb *sqlx.DB) error {
	err := db.StatusCheck(ctx, sqldb)
	if err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	driver, err := generic.New(sqldb.DB, postgres.Dialect{})
	if err != nil {
		return fmt.Errorf("construct darwin driver: %w", err)
	}

	d := darwin.New(driver, darwin.ParseMigrations(migrateDoc))
	return d.Migrate()
}

func Seed(ctx context.Context, sqldb *sqlx.DB) (err error) {
	err = db.StatusCheck(ctx, sqldb)
	if err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	tx, err := sqldb.Begin()
	if err != nil {
		return err
	}

	defer func() {
		errTx := tx.Rollback()
		if errTx != nil {
			if errors.Is(errTx, sql.ErrTxDone) {
				return
			}
			err = fmt.Errorf("rollback: %w", errTx)
			return
		}
	}()

	_, err = tx.Exec(seedDoc)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit: %w", err)
	}

	return nil
}
