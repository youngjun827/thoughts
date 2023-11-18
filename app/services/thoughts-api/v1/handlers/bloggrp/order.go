package bloggrp

import (
	"errors"
	"net/http"

	"github.com/youngjun827/thoughts/business/core/blog"
	"github.com/youngjun827/thoughts/business/database/order"
	"github.com/youngjun827/thoughts/foundation/validate"
)

func parseOrder(r *http.Request) (order.By, error) {
	const (
		orderByPostID      = "post_id"
		orderByTitle    = "title"
		orderByCategory   = "category"
		orderByEnabled = "enabled"
	)

	var orderByFields = map[string]string{
		orderByPostID:      blog.OrderByPostID,
		orderByTitle:    blog.OrderByTitle,
		orderByCategory:   blog.OrderByCategory,
		orderByEnabled: blog.OrderByEnabled,
	}

	orderBy, err := order.Parse(r, order.NewBy(orderByPostID, order.ASC))
	if err != nil {
		return order.By{}, err
	}

	_, exists := orderByFields[orderBy.Field]
	if !exists {
		return order.By{}, validate.NewFieldsError(orderBy.Field, errors.New("order field does not exist"))
	}

	orderBy.Field = orderByFields[orderBy.Field]

	return orderBy, nil
}