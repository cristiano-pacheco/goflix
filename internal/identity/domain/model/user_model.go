package model

import (
	"crypto/subtle"
	"errors"
	"strings"
	"time"
)

const (
	minPasswordHashLength = 32
	maxPasswordHashLength = 255
	minTokenLength        = 16
	maxTokenLength        = 255
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
	// Trim whitespace from inputs
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	passwordHash = strings.TrimSpace(passwordHash)
	confirmationToken = strings.TrimSpace(confirmationToken)

	// Validate inputs
	if err := validateUserCreationInputs(passwordHash, confirmationToken, confirmationExpiresAt); err != nil {
		return UserModel{}, err
	}

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

	// Create user model
	return UserModel{
		name:                  nameModel,
		email:                 emailModel,
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
	// Trim whitespace from inputs
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	passwordHash = strings.TrimSpace(passwordHash)

	// Validate inputs
	if err := validateUserRestorationInputs(id, passwordHash, createdAt, updatedAt); err != nil {
		return UserModel{}, err
	}

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

	// Create user model with all fields
	return UserModel{
		id:                     id,
		name:                   nameModel,
		email:                  emailModel,
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
	// Check if confirmation is still pending
	if u.confirmationToken == nil || u.confirmationExpiresAt == nil || u.confirmedAt != nil {
		return false
	}

	// Check if token has expired
	if !u.confirmationExpiresAt.After(time.Now().UTC()) {
		return false
	}

	// Use constant-time comparison to prevent timing attacks
	return subtle.ConstantTimeCompare([]byte(*u.confirmationToken), []byte(token)) == 1
}

func (u *UserModel) SetResetPasswordDetails(token string, expiresAt time.Time) error {
	token = strings.TrimSpace(token)
	if err := validateResetPasswordToken(token, expiresAt); err != nil {
		return err
	}

	u.resetPasswordToken = &token
	u.resetPasswordExpiresAt = &expiresAt
	u.updatedAt = time.Now().UTC()
	return nil
}

func (u *UserModel) ClearResetPasswordDetails() {
	u.resetPasswordToken = nil
	u.resetPasswordExpiresAt = nil
	u.updatedAt = time.Now().UTC()
}

func (u *UserModel) IsResetPasswordTokenValid(token string) bool {
	// Check if reset password token exists
	if u.resetPasswordToken == nil || u.resetPasswordExpiresAt == nil {
		return false
	}

	// Check if token has expired
	if !u.resetPasswordExpiresAt.After(time.Now().UTC()) {
		return false
	}

	// Use constant-time comparison to prevent timing attacks
	return subtle.ConstantTimeCompare([]byte(*u.resetPasswordToken), []byte(token)) == 1
}

func (u *UserModel) UpdatePasswordHash(newPasswordHash string) error {
	newPasswordHash = strings.TrimSpace(newPasswordHash)
	if err := validatePasswordHash(newPasswordHash); err != nil {
		return err
	}

	u.passwordHash = newPasswordHash
	u.updatedAt = time.Now().UTC()
	return nil
}

func validateUserCreationInputs(
	passwordHash, confirmationToken string,
	confirmationExpiresAt time.Time,
) error {
	if err := validatePasswordHash(passwordHash); err != nil {
		return err
	}

	if err := validateConfirmationToken(confirmationToken); err != nil {
		return err
	}

	if err := validateConfirmationExpiresAt(confirmationExpiresAt); err != nil {
		return err
	}

	return nil
}

func validateUserRestorationInputs(id uint64, passwordHash string, createdAt, updatedAt time.Time) error {
	if id == 0 {
		return errors.New("user ID is required and must be greater than zero")
	}

	if err := validatePasswordHash(passwordHash); err != nil {
		return err
	}

	if err := validateTimestamps(createdAt, updatedAt); err != nil {
		return err
	}

	return nil
}

func validatePasswordHash(passwordHash string) error {
	if len(passwordHash) == 0 {
		return errors.New("password hash is required")
	}

	// Basic validation for common hash formats (bcrypt, argon2, etc.)
	if len(passwordHash) < minPasswordHashLength {
		return errors.New("password hash appears to be too short (minimum 32 characters)")
	}

	if len(passwordHash) > maxPasswordHashLength {
		return errors.New("password hash exceeds maximum length of 255 characters")
	}

	return nil
}

func validateConfirmationToken(token string) error {
	if len(token) == 0 {
		return errors.New("confirmation token is required")
	}

	if len(token) < minTokenLength {
		return errors.New("confirmation token must be at least 16 characters long")
	}

	if len(token) > maxTokenLength {
		return errors.New("confirmation token exceeds maximum length of 255 characters")
	}

	return nil
}

func validateConfirmationExpiresAt(expiresAt time.Time) error {
	if expiresAt.IsZero() {
		return errors.New("confirmation expiration time is required")
	}

	if !expiresAt.After(time.Now().UTC()) {
		return errors.New("confirmation expiration time must be in the future")
	}

	return nil
}

func validateResetPasswordToken(token string, expiresAt time.Time) error {
	if len(token) == 0 {
		return errors.New("reset password token is required")
	}

	if len(token) < minTokenLength {
		return errors.New("reset password token must be at least 16 characters long")
	}

	if len(token) > maxTokenLength {
		return errors.New("reset password token exceeds maximum length of 255 characters")
	}

	if expiresAt.IsZero() {
		return errors.New("reset password expiration time is required")
	}

	if !expiresAt.After(time.Now().UTC()) {
		return errors.New("reset password expiration time must be in the future")
	}

	return nil
}

func validateTimestamps(createdAt, updatedAt time.Time) error {
	if createdAt.IsZero() {
		return errors.New("created at timestamp is required")
	}

	if updatedAt.IsZero() {
		return errors.New("updated at timestamp is required")
	}

	if updatedAt.Before(createdAt) {
		return errors.New("updated at timestamp cannot be before created at timestamp")
	}

	return nil
}
