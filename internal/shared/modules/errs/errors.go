package errs

import (
	"errors"
	"net/http"
)

var (
	// Authentication & Authorization
	ErrUserIsNotActivated              = errors.New("the user is not activated")
	ErrInvalidCredentials              = errors.New("invalid credentials")
	ErrInvalidToken                    = errors.New("invalid token")
	ErrInvalidAccountConfirmationToken = errors.New("invalid account confirmation token")

	// Database
	ErrNotFound = errors.New("not found")

	// private key
	ErrKeyMustBePEMEncoded = errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")

	// Bad Request
	ErrBadRequest = errors.New("bad request")
)

func NewBadRequestError(message string) error {
	return &Error{
		Status:        http.StatusBadRequest,
		OriginalError: ErrBadRequest,
		Err: er{
			Code:    codeBadRequest,
			Message: message,
		},
	}
}
