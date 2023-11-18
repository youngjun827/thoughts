// Package blogdb provides a sqlx queries support for bloggrp.
package blogdb

import (
	"bytes"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/youngjun827/thoughts/business/core/blog"
	db "github.com/youngjun827/thoughts/business/database/dbsql/pgx"
	"github.com/youngjun827/thoughts/business/database/order"
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

	err := db.NamedExecContext(ctx, s.log, s.db, q, toDBBlog(blog))
	if err != nil {
		s.log.Error(ctx, "blogdb.Create", "error", err)
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// Query retrieves a list of existing users from the database.
func (s *Store) Query(ctx context.Context, filter blog.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]blog.Blog, error) {
	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
	SELECT
		post_id, title, content, category, enabled, date_created, date_updated
	FROM
		blog_posts`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbBlgs []dbBlog
	if err := db.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbBlgs); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	blgs, err := toCoreBlogSlice(dbBlgs)
	if err != nil {
		return nil, err
	}

	return blgs, nil
}

// Count returns the total number of users in the DB.
func (s *Store) Count(ctx context.Context, filter blog.QueryFilter) (int, error) {
	data := map[string]interface{}{}

	const q = `
	SELECT
		count(1)
	FROM
		blog_posts`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	var count struct {
		Count int `db:"count"`
	}
	if err := db.NamedQueryStruct(ctx, s.log, s.db, buf.String(), data, &count); err != nil {
		return 0, fmt.Errorf("namedquerystruct: %w", err)
	}

	return count.Count, nil
}