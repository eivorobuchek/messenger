package models

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrAlreadyRegister = errors.New("already register")
)
