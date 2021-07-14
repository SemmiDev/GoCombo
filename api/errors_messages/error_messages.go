package errors_messages

import "errors"

var (
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrEmailAlreadyTaken    = errors.New("email already taken")

	ErrIncorrectPassword = errors.New("incorrect password")
	ErrIncorrectDetails  = errors.New("incorrect details")

	ErrUsernameRequired = errors.New("username required")
	ErrNameRequired     = errors.New("name required")
	ErrFullNameRequired = errors.New("full name required")
	ErrPasswordRequired = errors.New("password required")
	ErrEmailRequired    = errors.New("email required")

	ErrInvalidEmail = errors.New("invalid email")

	ErrUserNotFound     = errors.New("user not found")
	ErrProvinceNotFound = errors.New("province not found")
)
