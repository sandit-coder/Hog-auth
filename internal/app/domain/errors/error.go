package errors

import "errors"

var (
	AlreadyExists = errors.New("already exists")
	NotFound      = errors.New("not found")
	InvalidInput  = errors.New("invalid input")
)
