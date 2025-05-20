package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateLoginTokenModel(t *testing.T) {
	tests := []struct {
		name          string
		userID        uint64
		token         string
		expiresAt     time.Time
		expectError   bool
		errorContains string
	}{
		{
			name:        "Valid login token",
			userID:      123,
			token:       "valid_token",
			expiresAt:   time.Now().Add(24 * time.Hour),
			expectError: false,
		},
		{
			name:          "Zero user ID",
			userID:        0,
			token:         "valid_token",
			expiresAt:     time.Now().Add(24 * time.Hour),
			expectError:   true,
			errorContains: "user ID is required",
		},
		{
			name:          "Empty token",
			userID:        123,
			token:         "",
			expiresAt:     time.Now().Add(24 * time.Hour),
			expectError:   true,
			errorContains: "token is required",
		},
		{
			name:          "Zero expiration time",
			userID:        123,
			token:         "valid_token",
			expiresAt:     time.Time{},
			expectError:   true,
			errorContains: "expiration time is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := CreateLoginTokenModel(tt.userID, tt.token, tt.expiresAt)

			if tt.expectError {
				assert.Error(t, err)
				if err != nil {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				require.NotNil(t, token)

				// Verify fields
				assert.Equal(t, tt.userID, token.UserID())
				assert.Equal(t, tt.token, token.Token())
				assert.Equal(t, tt.expiresAt.Unix(), token.ExpiresAt().Unix())

				// Verify timestamps are set
				assert.False(t, token.CreatedAt().IsZero())
				assert.False(t, token.UpdatedAt().IsZero())

				// CreatedAt and UpdatedAt should be very close to each other
				timeDiff := token.UpdatedAt().Sub(token.CreatedAt())
				assert.LessOrEqual(t, timeDiff.Milliseconds(), int64(100))
			}
		})
	}
}

func TestRestoreLoginTokenModel(t *testing.T) {
	now := time.Now().UTC()
	expiresAt := now.Add(24 * time.Hour)

	tests := []struct {
		name          string
		id            uint64
		userID        uint64
		token         string
		expiresAt     time.Time
		createdAt     time.Time
		updatedAt     time.Time
		expectError   bool
		errorContains string
	}{
		{
			name:        "Valid login token with consumed time",
			id:          123,
			userID:      456,
			token:       "valid_token",
			expiresAt:   expiresAt,
			createdAt:   now.Add(-2 * time.Hour),
			updatedAt:   now,
			expectError: false,
		},
		{
			name:        "Valid login token without consumed time",
			id:          123,
			userID:      456,
			token:       "valid_token",
			expiresAt:   expiresAt,
			createdAt:   now.Add(-2 * time.Hour),
			updatedAt:   now,
			expectError: false,
		},
		{
			name:          "Zero ID",
			id:            0,
			userID:        456,
			token:         "valid_token",
			expiresAt:     expiresAt,
			createdAt:     now.Add(-2 * time.Hour),
			updatedAt:     now,
			expectError:   true,
			errorContains: "ID is required",
		},
		{
			name:          "Zero user ID",
			id:            123,
			userID:        0,
			token:         "valid_token",
			expiresAt:     expiresAt,
			createdAt:     now.Add(-2 * time.Hour),
			updatedAt:     now,
			expectError:   true,
			errorContains: "user ID is required",
		},
		{
			name:          "Empty token",
			id:            123,
			userID:        456,
			token:         "",
			expiresAt:     expiresAt,
			createdAt:     now.Add(-2 * time.Hour),
			updatedAt:     now,
			expectError:   true,
			errorContains: "token is required",
		},
		{
			name:          "Zero expiration time",
			id:            123,
			userID:        456,
			token:         "valid_token",
			expiresAt:     time.Time{},
			createdAt:     now.Add(-2 * time.Hour),
			updatedAt:     now,
			expectError:   true,
			errorContains: "expiration time is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := RestoreLoginTokenModel(
				tt.id,
				tt.userID,
				tt.token,
				tt.expiresAt,
				tt.createdAt,
				tt.updatedAt,
			)

			if tt.expectError {
				assert.Error(t, err)
				if err != nil {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				require.NotNil(t, token)

				// Verify fields
				assert.Equal(t, tt.id, token.ID())
				assert.Equal(t, tt.userID, token.UserID())
				assert.Equal(t, tt.token, token.Token())
				assert.Equal(t, tt.expiresAt.Unix(), token.ExpiresAt().Unix())

				assert.Equal(t, tt.createdAt.Unix(), token.CreatedAt().Unix())
				assert.Equal(t, tt.updatedAt.Unix(), token.UpdatedAt().Unix())
			}
		})
	}
}

func TestLoginTokenModel_IsExpired(t *testing.T) {
	now := time.Now().UTC()

	// Test case 1: Token not expired
	expiresAt := now.Add(24 * time.Hour)
	token, err := CreateLoginTokenModel(123, "valid_token", expiresAt)
	require.NoError(t, err)
	assert.False(t, token.IsExpired())

	// Test case 2: Token expired
	expiredExpiresAt := now.Add(-1 * time.Hour)
	expiredToken, err := RestoreLoginTokenModel(
		123,
		456,
		"expired_token",
		expiredExpiresAt,
		now.Add(-2*time.Hour),
		now,
	)
	require.NoError(t, err)
	assert.True(t, expiredToken.IsExpired())
}

func TestLoginTokenModel_IsValid(t *testing.T) {
	now := time.Now().UTC()

	testCases := []struct {
		name      string
		expiresAt time.Time
		expected  bool
	}{
		{
			name:      "Valid token (not expired, not consumed)",
			expiresAt: now.Add(24 * time.Hour),
			expected:  true,
		},
		{
			name:      "Invalid token (expired, not consumed)",
			expiresAt: now.Add(-1 * time.Hour),
			expected:  false,
		},
		{
			name:      "Invalid token (not expired, consumed)",
			expiresAt: now.Add(24 * time.Hour),
			expected:  true,
		},
		{
			name:      "Invalid token (expired, consumed)",
			expiresAt: now.Add(-1 * time.Hour),
			expected:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := RestoreLoginTokenModel(
				123,
				456,
				"test_token",
				tc.expiresAt,
				now.Add(-2*time.Hour),
				now,
			)
			require.NoError(t, err)

			assert.Equal(t, tc.expected, token.IsValid())
		})
	}
}
