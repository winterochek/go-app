package apiserver

import (
	"database/sql"
	"net/http"

	sqlstore "github.com/winterochek/go-app/internal/app/store/sql-store"
)

func Start(c *Config) error {
	db, err := NewDB(c.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	st := sqlstore.New(db)
	srv := NewServer(st)
	return http.ListenAndServe(c.BindAddr, srv)
}

func NewDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
