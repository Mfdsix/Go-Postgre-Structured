package database

import (
	"log"

	"go-postgre/app/models"

	"github.com/jmoiron/sqlx"
)

type PostDB interface {
	Open() error
	Close() error
	CreatePost(p *models.DPost) error
	GetPosts() ([]*models.DPost, error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", pgConnStr)

	if err != nil {
		return err
	}

	log.Print("Connected to Database!")
	pg.MustExec(CreateSchema)
	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
