package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winterochek/go-app/internal/app/model"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "with encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encrypted_password"
				return u
			},
			isValid: true,
		},
		{
			name: "invalid: empty field:email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid: invalid field:email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "not.email@"
				return u
			},
			isValid: false,
		},
		{
			name: "invalid: empty field:password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid: too short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "7chars_"
				return u
			},
			isValid: false,
		},
		{
			name: "invalid: too long password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "more_than_20_chars____"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
