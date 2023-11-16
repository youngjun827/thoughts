package blog

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/youngjun827/thoughts/foundation/validate"
)

type QueryFilter struct {
	PostID				*uuid.UUID
	Title				*string
	Content				*string
	Category			*string
	StartCreatedDate	*time.Time
	EndCreatedDate		*time.Time
}

func (qf *QueryFilter) Validate() error {
	err := validate.Check(qf)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

func (qf *QueryFilter) WithPostID(postID uuid.UUID) {
	qf.PostID = &postID
}

func (qf *QueryFilter) WithTitle(title string) {
	qf.Title = &title
}

func (qf *QueryFilter) WithContent(content string) {
	qf.Content = &content
}

func (qf *QueryFilter) WithCategory(category string) {
	qf.Category = &category
}

func (qf *QueryFilter) WithStartDateCreated(startDate time.Time) {
	d := startDate.UTC()
	qf.StartCreatedDate = &d
}

func (qf *QueryFilter) WithEndCreatedDate(endDate time.Time) {
	d := endDate.UTC()
	qf.EndCreatedDate = &d
}