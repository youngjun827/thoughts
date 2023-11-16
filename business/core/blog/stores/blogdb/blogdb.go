// Package blogdb provides a sqlx queries support for bloggrp.
package blogdb

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/youngjun827/thoughts/business/core/blog"
	db "github.com/youngjun827/thoughts/business/data/dbsql/pgx"
	"github.com/youngjun827/thoughts/foundation/logger"
)

type Store struct {
	log *logger.Logger
	db  *sqlx.DB
}

func NewStore(log *logger.Logger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

func (s *Store) Create(ctx context.Context, blog blog.Blog) error {
	const q = `
	INSERT INTO blog_posts
		(post_id, title, content, category, enabled, date_created, date_updated)
	VALUES
		(:post_id, :title, :content, :category, :enabled, :date_created, :date_updated)`

	fmt.Println(s.db.Rebind(q), toDBBlog(blog))

	err := db.NamedExecContext(ctx, s.log, s.db, q, toDBBlog(blog))
	if err != nil {
		s.log.Error(ctx, "blogdb.Create", "error", err)
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}
