package errs

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

// Password validation errors.
var (
	ErrPasswordTooShort      = errors.New("password must be at least 8 characters long")
	ErrPasswordNoUppercase   = errors.New("password must contain at least one uppercase letter")
	ErrPasswordNoLowercase   = errors.New("password must contain at least one lowercase letter")
	ErrPasswordNoNumber      = errors.New("password must contain at least one number")
	ErrPasswordNoSpecialChar = errors.New("password must contain at least one special character")
	ErrEmailAlreadyInUse     = errors.New("email already in use")
)
