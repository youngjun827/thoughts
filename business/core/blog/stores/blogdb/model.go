package blogdb

import (
	"time"

	"github.com/google/uuid"
	"github.com/youngjun827/thoughts/business/core/blog"
)

type dbBlog struct {
	PostID      uuid.UUID `db:"post_id"`
	Title       string    `db:"title"`
	Content     string    `db:"content"`
	Category    string    `db:"category"`
	Enabled     bool      `db:"enabled"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBBlog(blog blog.Blog) dbBlog {
	return dbBlog{
		PostID:      blog.PostID,
		Title:       blog.Title,
		Content:     blog.Content,
		Category:    blog.Category,
		Enabled:     blog.Enabled,
		DateCreated: blog.DateCreated.UTC(),
		DateUpdated: blog.DateUpdated.UTC(),
	}
}


func toCoreBlog(dbBlog dbBlog) (blog.Blog, error) {
	blog := blog.Blog{
		PostID:			dbBlog.PostID,
		Title:          dbBlog.Title,
		Content:        dbBlog.Content,
		Category:		dbBlog.Category,
		Enabled:        dbBlog.Enabled,
		DateCreated:    dbBlog.DateCreated.In(time.Local),
		DateUpdated:    dbBlog.DateUpdated.In(time.Local),
	}

	return blog, nil
}

func toCoreBlogSlice(dbBlogs []dbBlog) ([]blog.Blog, error) {
	blogs := make([]blog.Blog, len(dbBlogs))
	for i, dbBlog := range dbBlogs {
		var err error
		blogs[i], err = toCoreBlog(dbBlog)
		if err != nil {
			return nil, err
		}
	}
	return blogs, nil
}
