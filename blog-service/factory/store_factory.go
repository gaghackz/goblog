package factory

import (
	"blog-service/store"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gofr.dev/pkg/gofr"
)

func NewBlogStore(app *gofr.App) store.BlogStore {
	
	db, err := sql.Open("mysql", "root:password@localhost:2001/test")
	if err != nil {
		panic("failed to connect to database")
	}

	// Test connection
	if err = db.Ping(); err != nil {
		panic("cannot connect to dockerized MySQL")
	}

	return store.New(db)
}