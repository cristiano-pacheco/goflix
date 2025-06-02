package model

import (
	"time"

	"github.com/samber/lo"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/errs"
)

type AuthTokenModel struct {
	id        uint64
	userID    uint64
	token     string
	expiresAt time.Time
	createdAt time.Time
	updatedAt time.Time
}

func CreateAuthTokenModel(userID uint64, token string, expiresAt time.Time) (AuthTokenModel, error) {
	if userID == 0 {
		return AuthTokenModel{}, errs.ErrAuthTokenUserIDRequired
	}

	if lo.IsEmpty(token) {
		return AuthTokenModel{}, errs.ErrAuthTokenTokenRequired
	}

	if expiresAt.IsZero() {
		return AuthTokenModel{}, errs.ErrAuthTokenExpirationRequired
	}

	return AuthTokenModel{
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}, nil
}

func RestoreAuthTokenModel(
	id uint64,
	userID uint64,
	token string,
	expiresAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (AuthTokenModel, error) {
	if id == 0 {
		return AuthTokenModel{}, errs.ErrAuthTokenIDRequired
	}

	if userID == 0 {
		return AuthTokenModel{}, errs.ErrAuthTokenUserIDRequired
	}

	if lo.IsEmpty(token) {
		return AuthTokenModel{}, errs.ErrAuthTokenTokenRequired
	}

	if expiresAt.IsZero() {
		return AuthTokenModel{}, errs.ErrAuthTokenExpirationRequired
	}

	return AuthTokenModel{
		id:        id,
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (t *AuthTokenModel) ID() uint64 {
	return t.id
}

func (t *AuthTokenModel) UserID() uint64 {
	return t.userID
}

func (t *AuthTokenModel) Token() string {
	return t.token
}

func (t *AuthTokenModel) ExpiresAt() time.Time {
	return t.expiresAt
}

func (t *AuthTokenModel) CreatedAt() time.Time {
	return t.createdAt
}

func (t *AuthTokenModel) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *AuthTokenModel) IsExpired() bool {
	return time.Now().UTC().After(t.expiresAt)
}

func (t *AuthTokenModel) IsValid() bool {
	return !t.IsExpired()
}
