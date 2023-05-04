package constants

import "errors"

const (
	CheckUsernameQuery = "SELECT id, username, role, password FROM users WHERE username = ?;"
)

var (
	ErrInvalidResources          = errors.New("invalid resources")
	ErrMismatchedHashAndPassword = errors.New("wrong password")
	ErrNoUsernameExist           = errors.New("username not found")
)

var (
	InternalServerError = "Internal Server Error"
)

var (
	LoginSuccess = "Login Success!"
)
