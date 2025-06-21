package store

import (
	"blog-service/model"
	"database/sql"

	"gofr.dev/pkg/gofr"
)

type BlogStore interface {
	Create(*gofr.Context, model.Blog) (model.Blog, error)
	GetAll(*gofr.Context) ([]model.Blog, error)
	GetByID(*gofr.Context, string) (model.Blog, error)
}

type blogStore struct {
	db *sql.DB
}

func New(db *sql.DB) BlogStore {
	return &blogStore{db: db}
}

func (s *blogStore) Create(ctx *gofr.Context, b model.Blog) (model.Blog, error) {
	_, err := s.db.ExecContext(ctx, "INSERT INTO blogs (id, title, content, author) VALUES (?, ?, ?, ?)", b.ID, b.Title, b.Content, b.Author)
	return b, err
}

func (s *blogStore) GetAll(ctx *gofr.Context) ([]model.Blog, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, title, content, author FROM blogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []model.Blog
	for rows.Next() {
		var b model.Blog
		err := rows.Scan(&b.ID, &b.Title, &b.Content, &b.Author)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}
	return blogs, nil
}

func (s *blogStore) GetByID(ctx *gofr.Context, id string) (model.Blog, error) {
	var b model.Blog
	err := s.db.QueryRowContext(ctx, "SELECT id, title, content, author FROM blogs WHERE id = ?", id).Scan(&b.ID, &b.Title, &b.Content, &b.Author)
	return b, err
}