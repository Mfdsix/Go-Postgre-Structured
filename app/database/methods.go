package database

import "go-postgre/app/models"

func (d *DB) CreatePost(p *models.DPost) error {
	res, err := d.db.Exec(CreatePostSchema, p.Title, p.Content, p.Author)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) GetPosts() ([]*models.DPost, error) {
	var posts []*models.DPost
	err := d.db.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		return posts, err
	}

	return posts, nil
}
