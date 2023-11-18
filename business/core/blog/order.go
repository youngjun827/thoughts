package blog

import "github.com/youngjun827/thoughts/business/database/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByPostID, order.ASC)

// Set of fields that the results can be ordered by. These are the names
// that should be used by the application layer.
const (
	OrderByPostID      = "post_id"
	OrderByTitle    = "title"
	OrderByCategory   = "category"
	OrderByEnabled = "enabled"
)