package blog

import "github.com/youngjun827/thoughts/business/database/order"

var DefaultOrderBy = order.NewBy(OrderByPostID, order.ASC)

const (
	OrderByPostID   = "post_id"
	OrderByTitle    = "title"
	OrderByCategory = "category"
	OrderByEnabled  = "enabled"
)
