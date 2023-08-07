package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
	ur     *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	fmt.Println("db connected")
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.ur != nil {
		return s.ur
	}
	s.ur = &UserRepository{
		store: s,
	}
	return s.ur
}
