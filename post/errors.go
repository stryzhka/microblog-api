package post

import "errors"

var (
	ErrPostNotFound = errors.New("post not found")
	ErrAlreadyLiked = errors.New("already liked")
	ErrNotLiked     = errors.New("not liked yet")
)
