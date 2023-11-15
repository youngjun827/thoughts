INSERT INTO blog_posts (post_id, title, content, category_id, enabled, date_created, date_updated) 
VALUES
    ('1a4302e9-97a1-4e8a-ba02-0c5e36e8f212', 'Introduction to SQL', '...', '2f8cfb27-6a53-4f9a-8f45-9a69c50327c7', true, '2023-01-01 08:00:00', '2023-01-01 08:00:00')
ON CONFLICT (post_id) DO NOTHING;
