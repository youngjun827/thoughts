// Package bloggrp adds specific routes for CRUD operations related to blogs.
package bloggrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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


func (h *Handlers) QueryByPostID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	postIDStr := chi.URLParam(r, "post_id")
	if postIDStr == "" {
		return response.NewError(errors.New("post_id not found in URL"), http.StatusBadRequest)
	}

	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		return response.NewError(errors.New("invalid post_id"), http.StatusBadRequest)
	}

	blg, err := h.blog.QueryByPostID(ctx, postID)
	if err != nil {
		switch {
		case errors.Is(err, blog.ErrNotFound):
			return response.NewError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("querybyid: id[%s]: %w", postID, err)
		}
	}

	return web.Respond(ctx, w, toAppBlog(blg), http.StatusOK)
}