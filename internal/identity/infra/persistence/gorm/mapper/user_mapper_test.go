package mapper_test

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/entity"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/mapper"
)

func TestUserMapper_ToModel(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	confirmToken := "confirm_token"
	resetToken := "reset_token"
	confirmExpiry := now.Add(24 * time.Hour)
	resetExpiry := now.Add(12 * time.Hour)
	confirmedAt := now.Add(-1 * time.Hour)

	// Create a user entity
	userEntity := entity.UserEntity{
		ID:                     123,
		Name:                   "John Doe",
		Email:                  "john@example.com",
		PasswordHash:           "$2a$10$abcdefghijklmnopqrstuvwxyz123456789",
		IsActivated:            true,
		ConfirmationToken:      &confirmToken,
		ConfirmationExpiresAt:  &confirmExpiry,
		ConfirmedAt:            &confirmedAt,
		ResetPasswordToken:     &resetToken,
		ResetPasswordExpiresAt: &resetExpiry,
		CreatedAt:              now,
		UpdatedAt:              now,
	}

	// Create mapper
	mapper := mapper.NewUserMapper()

	// Test ToModel
	userModel, err := mapper.ToModel(userEntity)
	require.NoError(t, err)

	// Verify fields
	assert.Equal(t, uint64(123), userModel.ID())
	assert.Equal(t, "John Doe", userModel.Name())
	assert.Equal(t, "john@example.com", userModel.Email())
	assert.Equal(t, "$2a$10$abcdefghijklmnopqrstuvwxyz123456789", userModel.PasswordHash())
	assert.True(t, userModel.IsActivated())

	require.NotNil(t, userModel.ConfirmationToken())
	assert.Equal(t, confirmToken, *userModel.ConfirmationToken())

	require.NotNil(t, userModel.ConfirmationExpiresAt())
	assert.Equal(t, confirmExpiry.Unix(), userModel.ConfirmationExpiresAt().Unix())

	require.NotNil(t, userModel.ConfirmedAt())
	assert.Equal(t, confirmedAt.Unix(), userModel.ConfirmedAt().Unix())

	require.NotNil(t, userModel.ResetPasswordToken())
	assert.Equal(t, resetToken, *userModel.ResetPasswordToken())

	require.NotNil(t, userModel.ResetPasswordExpiresAt())
	assert.Equal(t, resetExpiry.Unix(), userModel.ResetPasswordExpiresAt().Unix())

	assert.Equal(t, now.Unix(), userModel.CreatedAt().Unix())
	assert.Equal(t, now.Unix(), userModel.UpdatedAt().Unix())
}

func TestUserMapper_ToEntity(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	confirmToken := "confirm_token"
	resetToken := "reset_token"
	confirmExpiry := now.Add(24 * time.Hour)
	resetExpiry := now.Add(12 * time.Hour)
	confirmedAt := now.Add(-1 * time.Hour)

	// Create a user model
	userModel, err := model.RestoreUserModel(
		123,
		"John Doe",
		"john@example.com",
		"$2a$10$abcdefghijklmnopqrstuvwxyz123456789",
		true,
		lo.ToPtr(confirmToken),
		lo.ToPtr(confirmExpiry),
		lo.ToPtr(confirmedAt),
		lo.ToPtr(resetToken),
		lo.ToPtr(resetExpiry),
		now,
		now,
	)
	require.NoError(t, err)

	// Create mapper
	mapper := mapper.NewUserMapper()

	// Test ToEntity
	userEntity := mapper.ToEntity(userModel)

	// Verify fields
	assert.Equal(t, uint64(123), userEntity.ID)
	assert.Equal(t, "John Doe", userEntity.Name)
	assert.Equal(t, "john@example.com", userEntity.Email)
	assert.Equal(t, "$2a$10$abcdefghijklmnopqrstuvwxyz123456789", userEntity.PasswordHash)
	assert.True(t, userEntity.IsActivated)

	require.NotNil(t, userEntity.ConfirmationToken)
	assert.Equal(t, confirmToken, *userEntity.ConfirmationToken)

	require.NotNil(t, userEntity.ConfirmationExpiresAt)
	assert.Equal(t, confirmExpiry.Unix(), userEntity.ConfirmationExpiresAt.Unix())

	require.NotNil(t, userEntity.ConfirmedAt)
	assert.Equal(t, confirmedAt.Unix(), userEntity.ConfirmedAt.Unix())

	require.NotNil(t, userEntity.ResetPasswordToken)
	assert.Equal(t, resetToken, *userEntity.ResetPasswordToken)

	require.NotNil(t, userEntity.ResetPasswordExpiresAt)
	assert.Equal(t, resetExpiry.Unix(), userEntity.ResetPasswordExpiresAt.Unix())

	assert.Equal(t, now.Unix(), userEntity.CreatedAt.Unix())
	assert.Equal(t, now.Unix(), userEntity.UpdatedAt.Unix())
}

func TestUserMapper_ToModel_WithNilPointers(t *testing.T) {
	// Create test data with nil pointers
	now := time.Now().UTC()

	// Create a user entity with nil pointers
	userEntity := entity.UserEntity{
		ID:                     123,
		Name:                   "John Doe",
		Email:                  "john@example.com",
		PasswordHash:           "$2a$10$abcdefghijklmnopqrstuvwxyz123456789",
		IsActivated:            true,
		ConfirmationToken:      nil,
		ConfirmationExpiresAt:  nil,
		ConfirmedAt:            nil,
		ResetPasswordToken:     nil,
		ResetPasswordExpiresAt: nil,
		CreatedAt:              now,
		UpdatedAt:              now,
	}

	// Create mapper
	mapper := mapper.NewUserMapper()

	// Test ToModel
	userModel, err := mapper.ToModel(userEntity)
	require.NoError(t, err)

	// Verify fields
	assert.Equal(t, uint64(123), userModel.ID())
	assert.Equal(t, "John Doe", userModel.Name())
	assert.Equal(t, "john@example.com", userModel.Email())
	assert.Equal(t, "$2a$10$abcdefghijklmnopqrstuvwxyz123456789", userModel.PasswordHash())
	assert.True(t, userModel.IsActivated())

	assert.Nil(t, userModel.ConfirmationToken())
	assert.Nil(t, userModel.ConfirmationExpiresAt())
	assert.Nil(t, userModel.ConfirmedAt())
	assert.Nil(t, userModel.ResetPasswordToken())
	assert.Nil(t, userModel.ResetPasswordExpiresAt())

	assert.Equal(t, now.Unix(), userModel.CreatedAt().Unix())
	assert.Equal(t, now.Unix(), userModel.UpdatedAt().Unix())
}

func TestUserMapper_ToEntity_WithNilPointers(t *testing.T) {
	// Create test data with nil pointers
	now := time.Now().UTC()

	// Create a user model with nil pointers
	userModel, err := model.RestoreUserModel(
		123,
		"John Doe",
		"john@example.com",
		"$2a$10$abcdefghijklmnopqrstuvwxyz123456789",
		true,
		nil,
		nil,
		nil,
		nil,
		nil,
		now,
		now,
	)
	require.NoError(t, err)

	// Create mapper
	mapper := mapper.NewUserMapper()

	// Test ToEntity
	userEntity := mapper.ToEntity(userModel)

	// Verify fields
	assert.Equal(t, uint64(123), userEntity.ID)
	assert.Equal(t, "John Doe", userEntity.Name)
	assert.Equal(t, "john@example.com", userEntity.Email)
	assert.Equal(t, "$2a$10$abcdefghijklmnopqrstuvwxyz123456789", userEntity.PasswordHash)
	assert.True(t, userEntity.IsActivated)

	assert.Nil(t, userEntity.ConfirmationToken)
	assert.Nil(t, userEntity.ConfirmationExpiresAt)
	assert.Nil(t, userEntity.ConfirmedAt)
	assert.Nil(t, userEntity.ResetPasswordToken)
	assert.Nil(t, userEntity.ResetPasswordExpiresAt)

	assert.Equal(t, now.Unix(), userEntity.CreatedAt.Unix())
	assert.Equal(t, now.Unix(), userEntity.UpdatedAt.Unix())
}

func TestUserMapper_InvalidModel(t *testing.T) {
	// Create test data with invalid values
	now := time.Now().UTC()

	// Create a user entity with invalid email
	userEntity := entity.UserEntity{
		ID:                     123,
		Name:                   "John Doe",
		Email:                  "invalid-email", // Invalid email
		PasswordHash:           "$2a$10$abcdefghijklmnopqrstuvwxyz123456789",
		IsActivated:            true,
		ConfirmationToken:      nil,
		ConfirmationExpiresAt:  nil,
		ConfirmedAt:            nil,
		ResetPasswordToken:     nil,
		ResetPasswordExpiresAt: nil,
		CreatedAt:              now,
		UpdatedAt:              now,
	}

	// Create mapper
	mapper := mapper.NewUserMapper()

	// Test ToModel with invalid data
	_, err := mapper.ToModel(userEntity)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid email format")
}
