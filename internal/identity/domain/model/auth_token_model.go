package model

import (
	"errors"
	"time"

	"github.com/samber/lo"
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
		return AuthTokenModel{}, errors.New("user ID is required")
	}

	if lo.IsEmpty(token) {
		return AuthTokenModel{}, errors.New("token is required")
	}

	if expiresAt.IsZero() {
		return AuthTokenModel{}, errors.New("expiration time is required")
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
		return AuthTokenModel{}, errors.New("ID is required")
	}

	if userID == 0 {
		return AuthTokenModel{}, errors.New("user ID is required")
	}

	if lo.IsEmpty(token) {
		return AuthTokenModel{}, errors.New("token is required")
	}

	if expiresAt.IsZero() {
		return AuthTokenModel{}, errors.New("expiration time is required")
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
