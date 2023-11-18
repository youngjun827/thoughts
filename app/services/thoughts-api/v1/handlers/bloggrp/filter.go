package bloggrp

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/youngjun827/thoughts/business/core/blog"
	"github.com/youngjun827/thoughts/foundation/validate"
)

func parseFilter(r *http.Request) (blog.QueryFilter, error) {
	const (
		filterByPostID           = "post_id"
		filterByTitle            = "title"
		filterByCategory         = "category"
		filterByStartCreatedDate = "start_created_date"
	)

	values := r.URL.Query()

	var filter blog.QueryFilter

	postID := values.Get(filterByPostID)
	if postID != "" {
		id, err := uuid.Parse(postID)
		if err != nil {
			return blog.QueryFilter{}, validate.NewFieldsError(filterByPostID, err)
		}
		filter.WithPostID(id)
	}

	title := values.Get(filterByTitle)
	if title != "" {
		filter.WithTitle(title)
	}

	category := values.Get(filterByCategory)
	if category != "" {
		filter.WithTitle(category)
	}

	createdDate := values.Get(filterByStartCreatedDate)
	if createdDate != "" {
		t, err := time.Parse(time.RFC3339, createdDate)
		if err != nil {
			return blog.QueryFilter{}, validate.NewFieldsError(filterByStartCreatedDate, err)
		}
		filter.WithStartDateCreated(t)
	}

	if err := filter.Validate(); err != nil {
		return blog.QueryFilter{}, err
	}

	return filter, nil
}
