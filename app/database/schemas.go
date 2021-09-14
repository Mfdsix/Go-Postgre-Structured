package database

const CreateSchema = `
CREATE TABLE IF NOT EXISTS posts
(
	id SERIAL PRIMARY KEY,
	title TEXT,
	content TEXT,
	author TEXT
)
`

const CreatePostSchema = `
INSERT INTO posts (title, content, author) VALUES ($1, $2, $3) RETURNING id
`
