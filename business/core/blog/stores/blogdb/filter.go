package blogdb

import (
	"bytes"
	"strings"

	"github.com/youngjun827/thoughts/business/core/blog"
)

func (s *Store) applyFilter(filter blog.QueryFilter, data map[string]interface{}, buf *bytes.Buffer) {
	var wc []string

	if filter.PostID != nil {
		data["post_id"] = *filter.PostID
		wc = append(wc, "post_id = :post_id")
	}

	if filter.Title != nil {
		data["title"] = (*filter.Title)
		wc = append(wc, "title = :title")
	}

	if filter.Category != nil {
		data["category"] = (*filter.Category)
		wc = append(wc, "category = :category")
	}

	if filter.StartCreatedDate != nil {
		data["start_date_created"] = *filter.StartCreatedDate
		wc = append(wc, "date_created >= :start_date_created")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}