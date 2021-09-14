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
