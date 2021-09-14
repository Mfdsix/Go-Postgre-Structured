package models

type DPost struct {
	ID      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Author  string `db:"author"`
}

type JPost struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type RPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
