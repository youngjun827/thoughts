INSERT INTO blog_posts (post_id, title, content, category, enabled, date_created, date_updated) 
VALUES
    ('1a4302e9-97a1-4e8a-ba02-0c5e36e8f212', 'Introduction to SQL', '...', 'Database', true, '2023-01-01 08:00:00', '2023-01-01 08:00:00')
ON CONFLICT (post_id) DO NOTHING;
