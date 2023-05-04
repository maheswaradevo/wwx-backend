package constants

import "errors"

var (
	ErrDataNotFound = errors.New("data not found")
)

var (
	ErrInvalidResources          = errors.New("invalid resources")
	ErrMismatchedHashAndPassword = errors.New("wrong password")
	ErrNoUsernameExist           = errors.New("username not found")
	InternalServerError          = "Internal Server Error"
	BindingRequestError          = "Error binding request"
)
