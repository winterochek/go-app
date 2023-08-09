package model

import (
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-ozzo/ozzo-validation/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
	Password          string `json:"password,omitempty"`
}

// Hash password before saving to DB
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

// Validation of user model
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(
			&u.Password,
			validation.By(validation.RuleFunc(requiredIf(u.EncryptedPassword == ""))),
			validation.Length(8, 20)),
	)
}

// Remove unwanted fields
func (u *User) Sanitaze() {
	u.Password = ""
}

func (u *User) ComparePasswords(p string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(p)) == nil
}

// Enctypt password with bcrypt lib
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
