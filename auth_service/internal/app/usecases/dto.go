package usecases

import (
	"auth_service/internal/app/models"
	"errors"
)

type RegisterUser struct {
	Email    models.Email
	Password string
}

func (ru RegisterUser) Validate() error {
	if ru.Email.String() == "" {
		return errors.New("email is required")
	}
	if ru.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type LoginUser struct {
	Email    models.Email
	Password string
}

func (lu LoginUser) Validate() error {
	if lu.Email == "" {
		return errors.New("email is required")
	}
	if lu.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
