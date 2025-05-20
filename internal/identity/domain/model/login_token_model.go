package model

import (
	"errors"
	"time"
)

type LoginTokenModel struct {
	id        uint64
	userID    uint64
	token     string
	expiresAt time.Time
	createdAt time.Time
	updatedAt time.Time
}

func CreateLoginTokenModel(userID uint64, token string, expiresAt time.Time) (LoginTokenModel, error) {
	if userID == 0 {
		return LoginTokenModel{}, errors.New("user ID is required")
	}

	if token == "" {
		return LoginTokenModel{}, errors.New("token is required")
	}

	if expiresAt.IsZero() {
		return LoginTokenModel{}, errors.New("expiration time is required")
	}

	return LoginTokenModel{
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}, nil
}

func RestoreLoginTokenModel(
	id uint64,
	userID uint64,
	token string,
	expiresAt time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (LoginTokenModel, error) {
	if id == 0 {
		return LoginTokenModel{}, errors.New("ID is required")
	}

	if userID == 0 {
		return LoginTokenModel{}, errors.New("user ID is required")
	}

	if token == "" {
		return LoginTokenModel{}, errors.New("token is required")
	}

	if expiresAt.IsZero() {
		return LoginTokenModel{}, errors.New("expiration time is required")
	}

	return LoginTokenModel{
		id:        id,
		userID:    userID,
		token:     token,
		expiresAt: expiresAt,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (t *LoginTokenModel) ID() uint64 {
	return t.id
}

func (t *LoginTokenModel) UserID() uint64 {
	return t.userID
}

func (t *LoginTokenModel) Token() string {
	return t.token
}

func (t *LoginTokenModel) ExpiresAt() time.Time {
	return t.expiresAt
}

func (t *LoginTokenModel) CreatedAt() time.Time {
	return t.createdAt
}

func (t *LoginTokenModel) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *LoginTokenModel) IsExpired() bool {
	return time.Now().UTC().After(t.expiresAt)
}

func (t *LoginTokenModel) IsValid() bool {
	return !t.IsExpired()
}
