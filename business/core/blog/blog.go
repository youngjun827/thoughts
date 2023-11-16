// Package blog defines a core blogging functionality
package blog

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/youngjun827/thoughts/foundation/logger"
)

type Storer interface{
	Create(ctx context.Context, blog Blog) error
}

// =============================================================================

type Core struct {
	storer Storer
	log *logger.Logger
}

func NewCore(log *logger.Logger, storer Storer) *Core {
	return &Core{
		storer: storer,
		log: log,
	}
}

func (c *Core) Create(ctx context.Context, nb NewBlog) (Blog, error) {
	now := time.Now()

	blog := Blog{
		PostID:			uuid.New(),
		Title:			nb.Title,
		Content:		nb.Content,
		Category:   	nb.Category,
		Enabled:    	true,
		DateCreated: 	now,
		DateUpdated: 	now,
	}

	err := c.storer.Create(ctx, blog)
	if err != nil {
		return Blog{}, fmt.Errorf("create: %w", err)
	}

	return blog, nil
}