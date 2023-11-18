// Package blog defines a core blogging functionality
package blog

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/youngjun827/thoughts/business/database/order"
	"github.com/youngjun827/thoughts/foundation/logger"
)

var (
	ErrNotFound              = errors.New("blog not found")
)

type Storer interface {
	Create(ctx context.Context, blog Blog) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Blog, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByPostID(ctx context.Context, postID uuid.UUID) (Blog, error)
}

// =============================================================================

type Core struct {
	storer Storer
	log    *logger.Logger
}

func NewCore(log *logger.Logger, storer Storer) *Core {
	return &Core{
		storer: storer,
		log:    log,
	}
}

func (c *Core) Create(ctx context.Context, nb NewBlog) (Blog, error) {
	now := time.Now()

	blog := Blog{
		PostID:      uuid.New(),
		Title:       nb.Title,
		Content:     nb.Content,
		Category:    nb.Category,
		Enabled:     true,
		DateCreated: now,
		DateUpdated: now,
	}

	err := c.storer.Create(ctx, blog)
	if err != nil {
		return Blog{}, fmt.Errorf("create: %w", err)
	}

	return blog, nil
}

func (c *Core) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]Blog, error) {
	users, err := c.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return users, nil
}

func (c *Core) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return c.storer.Count(ctx, filter)
}

func (c *Core) QueryByPostID(ctx context.Context, postID uuid.UUID) (Blog, error) {
	blog, err := c.storer.QueryByPostID(ctx, postID)
	if err != nil {
		return Blog{}, fmt.Errorf("query: userID[%s]: %w", postID, err)
	}

	return blog, nil
}