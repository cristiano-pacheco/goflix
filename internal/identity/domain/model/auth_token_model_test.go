package model_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
)

func TestCreateAuthTokenModel(t *testing.T) {
	t.Run("valid input returns auth token model", func(t *testing.T) {
		// Arrange
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(time.Hour)

		// Act
		result, err := model.CreateAuthTokenModel(userID, token, expiresAt)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, userID, result.UserID())
		assert.Equal(t, token, result.Token())
		assert.Equal(t, expiresAt, result.ExpiresAt())
		assert.False(t, result.CreatedAt().IsZero())
		assert.False(t, result.UpdatedAt().IsZero())
		assert.Equal(t, uint64(0), result.ID())
	})

	t.Run("zero user ID returns error", func(t *testing.T) {
		// Arrange
		userID := uint64(0)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(time.Hour)

		// Act
		result, err := model.CreateAuthTokenModel(userID, token, expiresAt)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "user ID is required", err.Error())
		assert.Equal(t, model.AuthTokenModel{}, result)
	})

	t.Run("empty token returns error", func(t *testing.T) {
		// Arrange
		userID := uint64(123)
		token := ""
		expiresAt := time.Now().UTC().Add(time.Hour)

		// Act
		result, err := model.CreateAuthTokenModel(userID, token, expiresAt)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "token is required", err.Error())
		assert.Equal(t, model.AuthTokenModel{}, result)
	})

	t.Run("zero expiration time returns error", func(t *testing.T) {
		// Arrange
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Time{}

		// Act
		result, err := model.CreateAuthTokenModel(userID, token, expiresAt)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "expiration time is required", err.Error())
		assert.Equal(t, model.AuthTokenModel{}, result)
	})
}

func TestRestoreAuthTokenModel(t *testing.T) {
	t.Run("valid input returns auth token model", func(t *testing.T) {
		// Arrange
		id := uint64(456)
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(time.Hour)
		createdAt := time.Now().UTC().Add(-time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		result, err := model.RestoreAuthTokenModel(id, userID, token, expiresAt, createdAt, updatedAt)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, id, result.ID())
		assert.Equal(t, userID, result.UserID())
		assert.Equal(t, token, result.Token())
		assert.Equal(t, expiresAt, result.ExpiresAt())
		assert.Equal(t, createdAt, result.CreatedAt())
		assert.Equal(t, updatedAt, result.UpdatedAt())
	})

	t.Run("zero ID returns error", func(t *testing.T) {
		// Arrange
		id := uint64(0)
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(time.Hour)
		createdAt := time.Now().UTC().Add(-time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		result, err := model.RestoreAuthTokenModel(id, userID, token, expiresAt, createdAt, updatedAt)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "ID is required", err.Error())
		assert.Equal(t, model.AuthTokenModel{}, result)
	})

	t.Run("zero user ID returns error", func(t *testing.T) {
		// Arrange
		id := uint64(456)
		userID := uint64(0)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(time.Hour)
		createdAt := time.Now().UTC().Add(-time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		result, err := model.RestoreAuthTokenModel(id, userID, token, expiresAt, createdAt, updatedAt)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "user ID is required", err.Error())
		assert.Equal(t, model.AuthTokenModel{}, result)
	})

	t.Run("empty token returns error", func(t *testing.T) {
		// Arrange
		id := uint64(456)
		userID := uint64(123)
		token := ""
		expiresAt := time.Now().UTC().Add(time.Hour)
		createdAt := time.Now().UTC().Add(-time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		result, err := model.RestoreAuthTokenModel(id, userID, token, expiresAt, createdAt, updatedAt)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "token is required", err.Error())
		assert.Equal(t, model.AuthTokenModel{}, result)
	})

	t.Run("zero expiration time returns error", func(t *testing.T) {
		// Arrange
		id := uint64(456)
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Time{}
		createdAt := time.Now().UTC().Add(-time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		result, err := model.RestoreAuthTokenModel(id, userID, token, expiresAt, createdAt, updatedAt)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "expiration time is required", err.Error())
		assert.Equal(t, model.AuthTokenModel{}, result)
	})
}

func TestAuthTokenModel_IsExpired(t *testing.T) {
	t.Run("token with future expiration is not expired", func(t *testing.T) {
		// Arrange
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(time.Hour)
		authToken, _ := model.CreateAuthTokenModel(userID, token, expiresAt)

		// Act
		result := authToken.IsExpired()

		// Assert
		assert.False(t, result)
	})

	t.Run("token with past expiration is expired", func(t *testing.T) {
		// Arrange
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(-time.Hour)
		authToken, _ := model.CreateAuthTokenModel(userID, token, expiresAt)

		// Act
		result := authToken.IsExpired()

		// Assert
		assert.True(t, result)
	})
}

func TestAuthTokenModel_IsValid(t *testing.T) {
	t.Run("token with future expiration is valid", func(t *testing.T) {
		// Arrange
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(time.Hour)
		authToken, _ := model.CreateAuthTokenModel(userID, token, expiresAt)

		// Act
		result := authToken.IsValid()

		// Assert
		assert.True(t, result)
	})

	t.Run("token with past expiration is not valid", func(t *testing.T) {
		// Arrange
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(-time.Hour)
		authToken, _ := model.CreateAuthTokenModel(userID, token, expiresAt)

		// Act
		result := authToken.IsValid()

		// Assert
		assert.False(t, result)
	})
}

func TestAuthTokenModel_Getters(t *testing.T) {
	t.Run("getters return correct values", func(t *testing.T) {
		// Arrange
		id := uint64(456)
		userID := uint64(123)
		token := "valid-token"
		expiresAt := time.Now().UTC().Add(time.Hour)
		createdAt := time.Now().UTC().Add(-time.Hour)
		updatedAt := time.Now().UTC()
		authToken, _ := model.RestoreAuthTokenModel(id, userID, token, expiresAt, createdAt, updatedAt)

		// Act & Assert
		assert.Equal(t, id, authToken.ID())
		assert.Equal(t, userID, authToken.UserID())
		assert.Equal(t, token, authToken.Token())
		assert.Equal(t, expiresAt, authToken.ExpiresAt())
		assert.Equal(t, createdAt, authToken.CreatedAt())
		assert.Equal(t, updatedAt, authToken.UpdatedAt())
	})
}
