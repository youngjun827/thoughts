-- Version: 1.01
-- Description: Create table blog_posts
CREATE TABLE blog_posts (
	post_id       UUID        NOT NULL,
	title         TEXT        NOT NULL,
	content       TEXT UNIQUE NOT NULL,
	category_id   UUID        NOT NULL,
    enabled       BOOLEAN     NOT NULL,
	date_created  TIMESTAMP   NOT NULL,
	date_updated  TIMESTAMP   NOT NULL,

	PRIMARY KEY (post_id)
);
