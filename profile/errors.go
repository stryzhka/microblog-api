package profile

import "errors"

var (
	ErrProfileNotFound   = errors.New("profile not found")
	ErrNameAlreadyExists = errors.New("name already exists")
)
