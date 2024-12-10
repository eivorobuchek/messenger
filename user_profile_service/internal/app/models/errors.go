package models

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyFriend = errors.New("already friend")
)
