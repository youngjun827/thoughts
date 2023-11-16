package bloggrp

import (
	"time"

	"github.com/youngjun827/thoughts/business/core/blog"
	"github.com/youngjun827/thoughts/foundation/validate"
)

type AppBlog struct {
	PostID          string   `json:"post_id"`
	Title         	string   `json:"title"`
	Content       	string   `json:"content"`
	Category        string 	 `json:"category"`
	Enabled    	    bool     `json:"enabled"`
	DateCreated  	string   `json:"dateCreated"`
	DateUpdated  	string   `json:"dateUpdated"`
}

func toAppBlog(blog blog.Blog) AppBlog {
	return AppBlog{
		PostID:       blog.PostID.String(),
		Title:        blog.Title,
		Content:      blog.Content,
		Category:     blog.Category,
		Enabled:      blog.Enabled,
		DateCreated:  blog.DateCreated.Format(time.RFC3339),
		DateUpdated:  blog.DateUpdated.Format(time.RFC3339),
	}
}

/*
func toAppBlogs(blogs []blog.Blog) []AppBlog {
	items := make([]AppBlog, len(blogs))
	for i, usr := range blogs {
		items[i] = toAppBlog(usr)
	}

	return items
}
*/

// =============================================================================

type AppNewBlog struct {
	Title         	string   `json:"title" validate:"required"`
	Content       	string   `json:"content" validate:"required"`
	Category        string 	 `json:"category" validate:"required"`
}

func toCoreNewBlog(app AppNewBlog) (blog.NewBlog, error) {
	blog := blog.NewBlog{
		Title:            app.Title,
		Content:          app.Content,
		Category:         app.Category,
	}

	return blog, nil
}

func (app AppNewBlog) Validate() error {
	err := validate.Check(app)
	if err != nil {
		return err
	}

	return nil
}