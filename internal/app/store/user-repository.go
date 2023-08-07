package store

import "github.com/winterochek/go-app/internal/app/model"

type UserRepository struct {
	store *Store
}

func (ur *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := ur.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := ur.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}
	return u, nil
}
