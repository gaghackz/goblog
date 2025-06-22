package factory

import (
	"blog-service/store"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"gofr.dev/pkg/gofr"
)

func NewBlogStore(app *gofr.App) store.BlogStore {
	
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:2001)/test")
	if err != nil {
		panic("failed to connect to database")
	}

	// Test connection
	if err = db.Ping(); err != nil {
		panic("cannot connect to dockerized MySQL")
	}

	schema := `
	CREATE TABLE IF NOT EXISTS blogs (
		id VARCHAR(50) PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL
	);`

	if _, err := db.Exec(schema); err != nil {
		panic("failed to create blogs table: " + err.Error())
	}

	return store.New(db)

}