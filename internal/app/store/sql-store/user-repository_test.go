package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winterochek/go-app/internal/app/model"
	"github.com/winterochek/go-app/internal/app/store"
	sqlstore "github.com/winterochek/go-app/internal/app/store/sql-store"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.ID)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
