package domain

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	Id           string
	Email        string
	Password     string
	PasswordHash string
	IsSuperuser  bool
	CreatedAt    time.Time
	Active       bool
}

func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(3, 255)),
		validation.Field(&u.Password, validation.Required, validation.Length(3, 99)),
	)

}

type UserRepository interface {
	GetAll() ([]User, error)
	GetOneById(string) (User, error)
	GetOneByEmail(string) (User, error)
	CreateOne(*User) error
	UpdateOneById(string, *User) error
	DeleteOneById(string) error
}
