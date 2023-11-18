package blog

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	PostID      uuid.UUID
	Title       string
	Content     string
	Category    string
	Enabled     bool
	DateCreated time.Time
	DateUpdated time.Time
}

type NewBlog struct {
	Title    string
	Content  string
	Category string
}

type UpdateUser struct {
	Title    *string
	Content  *string
	Category *string
	Enabled  *bool
}
