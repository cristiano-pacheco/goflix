package errs

import (
	"errors"
)

// Common error types
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrInternal     = errors.New("internal error")
)
