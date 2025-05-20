package model

import (
	"errors"
	"time"
)

type UserModel struct {
	id                     uint64
	name                   NameModel
	email                  EmailModel
	passwordHash           string
	isActivated            bool
	confirmationToken      *string
	confirmationExpiresAt  *time.Time
	confirmedAt            *time.Time
	resetPasswordToken     *string
	resetPasswordExpiresAt *time.Time
	createdAt              time.Time
	updatedAt              time.Time
}

func CreateUserModel(
	name string,
	email string,
	passwordHash string,
	confirmationToken string,
	confirmationExpiresAt time.Time,
) (UserModel, error) {
	// Validate and create name model
	nameModel, err := CreateNameModel(name)
	if err != nil {
		return UserModel{}, err
	}

	// Validate and create email model
	emailModel, err := CreateEmailModel(email)
	if err != nil {
		return UserModel{}, err
	}

	// Validate password hash
	if passwordHash == "" {
		return UserModel{}, errors.New("password hash is required")
	}

	// Create user model
	return UserModel{
		name:                  *nameModel,
		email:                 *emailModel,
		passwordHash:          passwordHash,
		isActivated:           false,
		confirmationToken:     &confirmationToken,
		confirmationExpiresAt: &confirmationExpiresAt,
		createdAt:             time.Now().UTC(),
		updatedAt:             time.Now().UTC(),
	}, nil
}

func RestoreUserModel(
	id uint64,
	name string,
	email string,
	passwordHash string,
	isActivated bool,
	confirmationToken *string,
	confirmationExpiresAt *time.Time,
	confirmedAt *time.Time,
	resetPasswordToken *string,
	resetPasswordExpiresAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (UserModel, error) {
	// Validate and create name model
	nameModel, err := CreateNameModel(name)
	if err != nil {
		return UserModel{}, err
	}

	// Validate and create email model
	emailModel, err := CreateEmailModel(email)
	if err != nil {
		return UserModel{}, err
	}

	// Validate password hash
	if passwordHash == "" {
		return UserModel{}, errors.New("password hash is required")
	}

	// Create user model with all fields
	return UserModel{
		id:                     id,
		name:                   *nameModel,
		email:                  *emailModel,
		passwordHash:           passwordHash,
		isActivated:            isActivated,
		confirmationToken:      confirmationToken,
		confirmationExpiresAt:  confirmationExpiresAt,
		confirmedAt:            confirmedAt,
		resetPasswordToken:     resetPasswordToken,
		resetPasswordExpiresAt: resetPasswordExpiresAt,
		createdAt:              createdAt,
		updatedAt:              updatedAt,
	}, nil
}

func (u *UserModel) ID() uint64 {
	return u.id
}

func (u *UserModel) Name() string {
	return u.name.String()
}

func (u *UserModel) Email() string {
	return u.email.String()
}

func (u *UserModel) PasswordHash() string {
	return u.passwordHash
}

func (u *UserModel) IsActivated() bool {
	return u.isActivated
}

func (u *UserModel) ConfirmationToken() *string {
	return u.confirmationToken
}

func (u *UserModel) ConfirmationExpiresAt() *time.Time {
	return u.confirmationExpiresAt
}

func (u *UserModel) ConfirmedAt() *time.Time {
	return u.confirmedAt
}

func (u *UserModel) ResetPasswordToken() *string {
	return u.resetPasswordToken
}

func (u *UserModel) ResetPasswordExpiresAt() *time.Time {
	return u.resetPasswordExpiresAt
}

func (u *UserModel) CreatedAt() time.Time {
	return u.createdAt
}

func (u *UserModel) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *UserModel) Activate() {
	u.isActivated = true
	u.updatedAt = time.Now().UTC()
}

func (u *UserModel) ConfirmAccount() {
	now := time.Now().UTC()
	u.isActivated = true
	u.confirmedAt = &now
	u.confirmationToken = nil
	u.confirmationExpiresAt = nil
	u.updatedAt = now
}

func (u *UserModel) IsConfirmationTokenValid(token string) bool {
	return u.confirmationToken != nil &&
		u.confirmationExpiresAt != nil &&
		u.confirmedAt == nil &&
		(*u.confirmationExpiresAt).After(time.Now().UTC()) &&
		*u.confirmationToken == token
}

func (u *UserModel) SetResetPasswordDetails(token string, expiresAt time.Time) {
	u.resetPasswordToken = &token
	u.resetPasswordExpiresAt = &expiresAt
	u.updatedAt = time.Now().UTC()
}

func (u *UserModel) ClearResetPasswordDetails() {
	u.resetPasswordToken = nil
	u.resetPasswordExpiresAt = nil
	u.updatedAt = time.Now().UTC()
}
