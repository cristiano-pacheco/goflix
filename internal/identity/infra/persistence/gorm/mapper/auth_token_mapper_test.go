package mapper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/entity"

	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/mapper"
)

func TestAuthTokenMapper_ToModel(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	expiresAt := now.Add(24 * time.Hour)

	// Create a auth token entity
	authTokenEntity := entity.AuthTokenEntity{
		ID:        123,
		UserID:    456,
		Token:     "token123",
		ExpiresAt: expiresAt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Create mapper
	mapper := mapper.NewAuthTokenMapper()

	// Test ToModel
	authTokenModel, err := mapper.ToModel(authTokenEntity)
	require.NoError(t, err)

	// Verify fields
	assert.Equal(t, uint64(123), authTokenModel.ID())
	assert.Equal(t, uint64(456), authTokenModel.UserID())
	assert.Equal(t, "token123", authTokenModel.Token())
	assert.Equal(t, expiresAt.Unix(), authTokenModel.ExpiresAt().Unix())

	assert.Equal(t, now.Unix(), authTokenModel.CreatedAt().Unix())
	assert.Equal(t, now.Unix(), authTokenModel.UpdatedAt().Unix())

	// Test derived properties
	assert.False(t, authTokenModel.IsExpired()) // Shouldn't be expired yet
	assert.True(t, authTokenModel.IsValid())    // Not valid because it's consumed
}

func TestLoginTokenMapper_ToModel_WithNilConsumedAt(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	expiresAt := now.Add(24 * time.Hour)

	// Create a login token entity with nil ConsumedAt
	authTokenEntity := entity.AuthTokenEntity{
		ID:        123,
		UserID:    456,
		Token:     "token123",
		ExpiresAt: expiresAt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Create mapper
	mapper := mapper.NewAuthTokenMapper()

	// Test ToModel
	authTokenModel, err := mapper.ToModel(authTokenEntity)
	require.NoError(t, err)

	// Verify fields
	assert.Equal(t, uint64(123), authTokenModel.ID())
	assert.Equal(t, uint64(456), authTokenModel.UserID())
	assert.Equal(t, "token123", authTokenModel.Token())
	assert.Equal(t, expiresAt.Unix(), authTokenModel.ExpiresAt().Unix())
	assert.Equal(t, now.Unix(), authTokenModel.CreatedAt().Unix())
	assert.Equal(t, now.Unix(), authTokenModel.UpdatedAt().Unix())

	// Test derived properties
	assert.False(t, authTokenModel.IsExpired()) // Shouldn't be expired yet
	assert.True(t, authTokenModel.IsValid())    // Valid because it's not consumed and not expired
}

func TestLoginTokenMapper_ToEntity(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	expiresAt := now.Add(24 * time.Hour)

	// Create a login token model
	authTokenModel, err := model.RestoreAuthTokenModel(
		123,
		456,
		"token123",
		expiresAt,
		now,
		now,
	)
	require.NoError(t, err)

	// Create mapper
	mapper := mapper.NewAuthTokenMapper()

	// Test ToEntity
	authTokenEntity := mapper.ToEntity(authTokenModel)

	// Verify fields
	assert.Equal(t, uint64(123), authTokenEntity.ID)
	assert.Equal(t, uint64(456), authTokenEntity.UserID)
	assert.Equal(t, "token123", authTokenEntity.Token)
	assert.Equal(t, expiresAt.Unix(), authTokenEntity.ExpiresAt.Unix())

	assert.Equal(t, now.Unix(), authTokenEntity.CreatedAt.Unix())
	assert.Equal(t, now.Unix(), authTokenEntity.UpdatedAt.Unix())
}

func TestLoginTokenMapper_InvalidData(t *testing.T) {
	// Create test data with invalid values
	now := time.Now().UTC()

	// Test cases with invalid data
	testCases := []struct {
		name        string
		entity      entity.AuthTokenEntity
		errorString string
	}{
		{
			name: "Zero ID",
			entity: entity.AuthTokenEntity{
				ID:        0, // Invalid ID
				UserID:    456,
				Token:     "token123",
				ExpiresAt: now.Add(24 * time.Hour),
				CreatedAt: now,
				UpdatedAt: now,
			},
			errorString: "ID is required",
		},
		{
			name: "Zero UserID",
			entity: entity.AuthTokenEntity{
				ID:        123,
				UserID:    0, // Invalid UserID
				Token:     "token123",
				ExpiresAt: now.Add(24 * time.Hour),
				CreatedAt: now,
				UpdatedAt: now,
			},
			errorString: "user ID is required",
		},
		{
			name: "Empty Token",
			entity: entity.AuthTokenEntity{
				ID:        123,
				UserID:    456,
				Token:     "", // Invalid Token
				ExpiresAt: now.Add(24 * time.Hour),
				CreatedAt: now,
				UpdatedAt: now,
			},
			errorString: "token is required",
		},
	}

	// Create mapper
	mapper := mapper.NewAuthTokenMapper()

	// Test ToModel with invalid data
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := mapper.ToModel(tc.entity)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.errorString)
		})
	}
}
