// Package bloggrp adds specific routes for CRUD operations related to blogs.
package bloggrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/youngjun827/thoughts/business/core/blog"
	"github.com/youngjun827/thoughts/business/web/v1/response"
	"github.com/youngjun827/thoughts/foundation/web"
)

type Handlers struct {
	blog *blog.Core
}

func New(blog *blog.Core) *Handlers {
	return &Handlers{
		blog: blog,
	}
}

func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewBlog
	if err := web.Decode(r, &app); err != nil {
		return response.NewError(err, http.StatusBadRequest)
	}

	nc, err := toCoreNewBlog(app)
	if err != nil {
		return response.NewError(err, http.StatusBadRequest)
	}

	blg, err := h.blog.Create(ctx, nc)
	if err != nil {
		return fmt.Errorf("create: blog[%+v]: %w", blg, err)
	}

	return web.Respond(ctx, w, toAppBlog(blg), http.StatusCreated)
}