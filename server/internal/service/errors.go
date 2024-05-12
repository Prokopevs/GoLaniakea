package service

import "errors"

var (
	ErrNoSuchPost = errors.New("post with such id doesn't exist")
)