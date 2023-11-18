// Package bloggrp adds specific routes for CRUD operations related to blogs.
package bloggrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/youngjun827/thoughts/business/core/blog"
	"github.com/youngjun827/thoughts/business/database/page"
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

	err := web.Decode(r, &app)
	if err != nil {
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


// Query returns a list of users with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	page, err := page.Parse(r)
	if err != nil {
		return err
	}

	filter, err := parseFilter(r)
	if err != nil {
		return err
	}

	orderBy, err := parseOrder(r)
	if err != nil {
		return err
	}

	blogs, err := h.blog.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	total, err := h.blog.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, response.NewPageDocument(toAppBlogs(blogs), total, page.Number, page.RowsPerPage), http.StatusOK)
}