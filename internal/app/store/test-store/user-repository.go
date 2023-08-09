package teststore

import (
	"github.com/winterochek/go-app/internal/app/model"
	"github.com/winterochek/go-app/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (ur *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	ur.users[u.Email] = u
	u.ID = len(ur.users)

	return nil
}

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := ur.users[email]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
