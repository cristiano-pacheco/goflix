package model_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
)

func TestCreateUserModel(t *testing.T) {
	t.Run("valid user creation", func(t *testing.T) {
		// Arrange
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy" // bcrypt hash
		confirmationToken := "abc123def456ghi789jkl012"
		confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, name, user.Name())
		assert.Equal(t, email, user.Email())
		assert.Equal(t, passwordHash, user.PasswordHash())
		assert.False(t, user.IsActivated())
		assert.NotNil(t, user.ConfirmationToken())
		assert.Equal(t, confirmationToken, *user.ConfirmationToken())
		assert.NotNil(t, user.ConfirmationExpiresAt())
		assert.Nil(t, user.ConfirmedAt())
		assert.Nil(t, user.ResetPasswordToken())
		assert.Nil(t, user.ResetPasswordExpiresAt())
		assert.False(t, user.CreatedAt().IsZero())
		assert.False(t, user.UpdatedAt().IsZero())
	})

	t.Run("valid user creation with whitespace trimming", func(t *testing.T) {
		// Arrange
		name := "  John Doe  "
		email := "  john.doe@example.com  "
		passwordHash := "  $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy  "
		confirmationToken := "  abc123def456ghi789jkl012  "
		confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, "John Doe", user.Name())
		assert.Equal(t, "john.doe@example.com", user.Email())
		assert.Equal(t, "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", user.PasswordHash())
		assert.Equal(t, "abc123def456ghi789jkl012", *user.ConfirmationToken())
	})

	t.Run("invalid name", func(t *testing.T) {
		// Arrange
		name := "J" // too short
		email := "john.doe@example.com"
		passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
		confirmationToken := "abc123def456ghi789jkl012"
		confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.Error(t, err)
		assert.Equal(t, model.UserModel{}, user)
	})

	t.Run("invalid email", func(t *testing.T) {
		// Arrange
		name := "John Doe"
		email := "invalid-email"
		passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
		confirmationToken := "abc123def456ghi789jkl012"
		confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.Error(t, err)
		require.Equal(t, model.UserModel{}, user)
	})

	t.Run("empty password hash", func(t *testing.T) {
		// Arrange
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := ""
		confirmationToken := "abc123def456ghi789jkl012"
		confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.Error(t, err)
		require.Equal(t, "password hash is required", err.Error())
		require.Equal(t, model.UserModel{}, user)
	})

	t.Run("password hash too short", func(t *testing.T) {
		// Arrange
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := "short"
		confirmationToken := "abc123def456ghi789jkl012"
		confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.Error(t, err)
		require.Equal(t, "password hash appears to be too short (minimum 32 characters)", err.Error())
		require.Equal(t, model.UserModel{}, user)
	})

	t.Run("password hash too long", func(t *testing.T) {
		// Arrange
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := strings.Repeat("a", 256)
		confirmationToken := "abc123def456ghi789jkl012"
		confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.Error(t, err)
		require.Equal(t, "password hash exceeds maximum length of 255 characters", err.Error())
		require.Equal(t, model.UserModel{}, user)
	})

	t.Run("empty confirmation token", func(t *testing.T) {
		// Arrange
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
		confirmationToken := ""
		confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.Error(t, err)
		require.Equal(t, "confirmation token is required", err.Error())
		require.Equal(t, model.UserModel{}, user)
	})

	t.Run("confirmation token too short", func(t *testing.T) {
		// Arrange
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
		confirmationToken := "short"
		confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.Error(t, err)
		require.Equal(t, "confirmation token must be at least 16 characters long", err.Error())
		require.Equal(t, model.UserModel{}, user)
	})

	t.Run("confirmation expires in the past", func(t *testing.T) {
		// Arrange
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
		confirmationToken := "abc123def456ghi789jkl012"
		confirmationExpiresAt := time.Now().UTC().Add(-1 * time.Hour) // past time

		// Act
		user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)

		// Assert
		require.Error(t, err)
		require.Equal(t, "confirmation expiration time must be in the future", err.Error())
		require.Equal(t, model.UserModel{}, user)
	})
}

func TestRestoreUserModel(t *testing.T) {
	t.Run("valid user restoration", func(t *testing.T) {
		// Arrange
		id := uint64(123)
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
		isActivated := true
		confirmationToken := (*string)(nil)
		confirmationExpiresAt := (*time.Time)(nil)
		confirmedAt := &time.Time{}
		resetPasswordToken := (*string)(nil)
		resetPasswordExpiresAt := (*time.Time)(nil)
		createdAt := time.Now().UTC().Add(-24 * time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		user, err := model.RestoreUserModel(
			id, name, email, passwordHash, isActivated,
			confirmationToken, confirmationExpiresAt, confirmedAt,
			resetPasswordToken, resetPasswordExpiresAt,
			createdAt, updatedAt,
		)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, id, user.ID())
		assert.Equal(t, name, user.Name())
		assert.Equal(t, email, user.Email())
		assert.Equal(t, passwordHash, user.PasswordHash())
		assert.Equal(t, isActivated, user.IsActivated())
		assert.Equal(t, confirmationToken, user.ConfirmationToken())
		assert.Equal(t, confirmationExpiresAt, user.ConfirmationExpiresAt())
		assert.Equal(t, confirmedAt, user.ConfirmedAt())
		assert.Equal(t, resetPasswordToken, user.ResetPasswordToken())
		assert.Equal(t, resetPasswordExpiresAt, user.ResetPasswordExpiresAt())
		assert.Equal(t, createdAt, user.CreatedAt())
		assert.Equal(t, updatedAt, user.UpdatedAt())
	})

	t.Run("zero ID", func(t *testing.T) {
		// Arrange
		id := uint64(0)
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
		createdAt := time.Now().UTC().Add(-24 * time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		user, err := model.RestoreUserModel(
			id, name, email, passwordHash, false,
			nil, nil, nil, nil, nil,
			createdAt, updatedAt,
		)

		// Assert
		require.Error(t, err)
		require.Equal(t, "user ID is required and must be greater than zero", err.Error())
		require.Equal(t, model.UserModel{}, user)
	})

	t.Run("updated at before created at", func(t *testing.T) {
		// Arrange
		id := uint64(123)
		name := "John Doe"
		email := "john.doe@example.com"
		passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
		createdAt := time.Now().UTC()
		updatedAt := time.Now().UTC().Add(-1 * time.Hour) // before created at

		// Act
		user, err := model.RestoreUserModel(
			id, name, email, passwordHash, false,
			nil, nil, nil, nil, nil,
			createdAt, updatedAt,
		)

		// Assert
		require.Error(t, err)
		require.Equal(t, "updated at timestamp cannot be before created at timestamp", err.Error())
		require.Equal(t, model.UserModel{}, user)
	})
}

func TestUserModel_BusinessMethods(t *testing.T) {
	t.Run("Activate", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		originalUpdatedAt := user.UpdatedAt()

		// Act
		time.Sleep(1 * time.Millisecond) // ensure time difference
		user.Activate()

		// Assert
		assert.True(t, user.IsActivated())
		assert.True(t, user.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("ConfirmAccount", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		originalUpdatedAt := user.UpdatedAt()

		// Act
		time.Sleep(1 * time.Millisecond) // ensure time difference
		user.ConfirmAccount()

		// Assert
		assert.True(t, user.IsActivated())
		assert.Nil(t, user.ConfirmationToken())
		assert.Nil(t, user.ConfirmationExpiresAt())
		assert.NotNil(t, user.ConfirmedAt())
		assert.True(t, user.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("IsConfirmationTokenValid - valid token", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		token := *user.ConfirmationToken()

		// Act
		isValid := user.IsConfirmationTokenValid(token)

		// Assert
		assert.True(t, isValid)
	})

	t.Run("IsConfirmationTokenValid - invalid token", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)

		// Act
		isValid := user.IsConfirmationTokenValid("wrong-token")

		// Assert
		assert.False(t, isValid)
	})

	t.Run("IsConfirmationTokenValid - already confirmed", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		token := *user.ConfirmationToken()
		user.ConfirmAccount()

		// Act
		isValid := user.IsConfirmationTokenValid(token)

		// Assert
		assert.False(t, isValid)
	})

	t.Run("SetResetPasswordDetails - valid", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		token := "reset123token456here789"
		expiresAt := time.Now().UTC().Add(1 * time.Hour)
		originalUpdatedAt := user.UpdatedAt()

		// Act
		time.Sleep(1 * time.Millisecond) // ensure time difference
		err := user.SetResetPasswordDetails(token, expiresAt)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, token, *user.ResetPasswordToken())
		assert.Equal(t, expiresAt, *user.ResetPasswordExpiresAt())
		assert.True(t, user.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("SetResetPasswordDetails - invalid token", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		token := "short"
		expiresAt := time.Now().UTC().Add(1 * time.Hour)

		// Act
		err := user.SetResetPasswordDetails(token, expiresAt)

		// Assert
		require.Error(t, err)
		require.Equal(t, "reset password token must be at least 16 characters long", err.Error())
	})

	t.Run("ClearResetPasswordDetails", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		token := "reset123token456here789"
		expiresAt := time.Now().UTC().Add(1 * time.Hour)
		_ = user.SetResetPasswordDetails(token, expiresAt)
		originalUpdatedAt := user.UpdatedAt()

		// Act
		time.Sleep(1 * time.Millisecond) // ensure time difference
		user.ClearResetPasswordDetails()

		// Assert
		assert.Nil(t, user.ResetPasswordToken())
		assert.Nil(t, user.ResetPasswordExpiresAt())
		assert.True(t, user.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("IsResetPasswordTokenValid - valid token", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		token := "reset123token456here789"
		expiresAt := time.Now().UTC().Add(1 * time.Hour)
		_ = user.SetResetPasswordDetails(token, expiresAt)

		// Act
		isValid := user.IsResetPasswordTokenValid(token)

		// Assert
		assert.True(t, isValid)
	})

	t.Run("IsResetPasswordTokenValid - invalid token", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		token := "reset123token456here789"
		expiresAt := time.Now().UTC().Add(1 * time.Hour)
		_ = user.SetResetPasswordDetails(token, expiresAt)

		// Act
		isValid := user.IsResetPasswordTokenValid("wrong-token")

		// Assert
		assert.False(t, isValid)
	})

	t.Run("IsResetPasswordTokenValid - expired token", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		token := "reset123token456here789"
		expiresAt := time.Now().UTC().Add(-1 * time.Hour) // expired
		_ = user.SetResetPasswordDetails(token, expiresAt)

		// Act
		isValid := user.IsResetPasswordTokenValid(token)

		// Assert
		assert.False(t, isValid)
	})

	t.Run("UpdatePasswordHash - valid", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		newPasswordHash := "$2a$10$NewHashHereWithProperLength123456789"
		originalUpdatedAt := user.UpdatedAt()

		// Act
		time.Sleep(1 * time.Millisecond) // ensure time difference
		err := user.UpdatePasswordHash(newPasswordHash)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, newPasswordHash, user.PasswordHash())
		assert.True(t, user.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("UpdatePasswordHash - invalid", func(t *testing.T) {
		// Arrange
		user := createValidUser(t)
		newPasswordHash := "short"

		// Act
		err := user.UpdatePasswordHash(newPasswordHash)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "password hash appears to be too short (minimum 32 characters)", err.Error())
	})
}

func createValidUser(t *testing.T) model.UserModel {
	t.Helper()

	name := "John Doe"
	email := "john.doe@example.com"
	passwordHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"
	confirmationToken := "abc123def456ghi789jkl012"
	confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

	user, err := model.CreateUserModel(name, email, passwordHash, confirmationToken, confirmationExpiresAt)
	require.NoError(t, err)

	return user
}
