package auth

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidToken      = errors.New("invalid token")
)
