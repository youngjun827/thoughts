package blogdb

import (
	"fmt"

	"github.com/youngjun827/thoughts/business/core/blog"
	"github.com/youngjun827/thoughts/business/database/order"
)

var orderByFields = map[string]string{
	blog.OrderByPostID:   "post_id",
	blog.OrderByTitle:    "title",
	blog.OrderByCategory: "category",
	blog.OrderByEnabled:  "enabled",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
