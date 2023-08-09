package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	sqlstore "github.com/winterochek/go-app/internal/app/store/sql-store"
)

func Start(c *Config) error {
	db, err := NewDB(c.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	st := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(c.SessionKey))
	srv := NewServer(st, sessionStore)
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
