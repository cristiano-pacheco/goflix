package model

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUserModel(t *testing.T) {
	tests := []struct {
		name          string
		userName      string
		email         string
		passwordHash  string
		expectError   bool
		errorContains string
	}{
		{
			name:         "Valid user",
			userName:     "John Doe",
			email:        "john@example.com",
			passwordHash: "hashed_password",
			expectError:  false,
		},
		{
			name:          "Empty name",
			userName:      "",
			email:         "john@example.com",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "name is required",
		},
		{
			name:          "Name too short",
			userName:      "J",
			email:         "john@example.com",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "name must be at least 2 characters long",
		},
		{
			name:          "Name too long",
			userName:      string(make([]byte, 256)),
			email:         "john@example.com",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "name cannot exceed 255 characters",
		},
		{
			name:          "Invalid email",
			userName:      "John Doe",
			email:         "invalid-email",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "invalid email format",
		},
		{
			name:          "Empty email",
			userName:      "John Doe",
			email:         "",
			passwordHash:  "hashed_password",
			expectError:   true,
			errorContains: "email is required",
		},
		{
			name:          "Empty password hash",
			userName:      "John Doe",
			email:         "john@example.com",
			passwordHash:  "",
			expectError:   true,
			errorContains: "password hash is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use default confirmation token and expiration time for testing
			confirmationToken := "default_token"
			confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)

			user, err := CreateUserModel(tt.userName, tt.email, tt.passwordHash, confirmationToken, confirmationExpiresAt)

			if tt.expectError {
				assert.Error(t, err)
				if err != nil {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
			} else {
				assert.NoError(t, err)
				require.NotNil(t, user)

				// Verify name
				nameModel, err := CreateNameModel(tt.userName)
				require.NoError(t, err)
				assert.Equal(t, nameModel.String(), user.Name())

				// Verify email
				emailModel, err := CreateEmailModel(tt.email)
				require.NoError(t, err)
				assert.Equal(t, emailModel.String(), user.Email())

				// Verify other fields
				assert.Equal(t, tt.passwordHash, user.PasswordHash())
				assert.False(t, user.IsActivated())
				require.NotNil(t, user.ConfirmationToken())
				assert.Equal(t, confirmationToken, *user.ConfirmationToken())
				require.NotNil(t, user.ConfirmationExpiresAt())
				assert.Equal(t, confirmationExpiresAt.Unix(), user.ConfirmationExpiresAt().Unix())
				assert.Nil(t, user.ConfirmedAt())
				assert.Nil(t, user.ResetPasswordToken())
				assert.Nil(t, user.ResetPasswordExpiresAt())

				// Verify timestamps are set
				assert.False(t, user.CreatedAt().IsZero())
				assert.False(t, user.UpdatedAt().IsZero())

				// CreatedAt and UpdatedAt should be very close to each other
				timeDiff := user.UpdatedAt().Sub(user.CreatedAt())
				assert.LessOrEqual(t, timeDiff.Milliseconds(), int64(100))
			}
		})
	}
}

func TestUserModel_Fields(t *testing.T) {
	// Create test data
	now := time.Now().UTC()
	nameModel, err := CreateNameModel("John Doe")
	require.NoError(t, err)
	emailModel, err := CreateEmailModel("john@example.com")
	require.NoError(t, err)

	confirmToken := "confirm_token"
	resetToken := "reset_token"
	confirmExpiry := now.Add(24 * time.Hour)
	resetExpiry := now.Add(12 * time.Hour)
	confirmedAt := now.Add(-1 * time.Hour)

	// Create a user model manually for field testing
	user := UserModel{
		id:                     123,
		name:                   *nameModel,
		email:                  *emailModel,
		passwordHash:           "hashed_password",
		isActivated:            true,
		confirmationToken:      &confirmToken,
		confirmationExpiresAt:  &confirmExpiry,
		confirmedAt:            &confirmedAt,
		resetPasswordToken:     &resetToken,
		resetPasswordExpiresAt: &resetExpiry,
		createdAt:              now,
		updatedAt:              now,
	}

	// Test field values
	assert.Equal(t, uint64(123), user.ID())
	assert.Equal(t, "John Doe", user.Name())
	assert.Equal(t, "john@example.com", user.Email())
	assert.Equal(t, "hashed_password", user.PasswordHash())
	assert.True(t, user.IsActivated())
	assert.Equal(t, confirmToken, *user.ConfirmationToken())
	assert.Equal(t, confirmExpiry.Unix(), user.ConfirmationExpiresAt().Unix())
	assert.Equal(t, confirmedAt.Unix(), user.ConfirmedAt().Unix())
	assert.Equal(t, resetToken, *user.ResetPasswordToken())
	assert.Equal(t, resetExpiry.Unix(), user.ResetPasswordExpiresAt().Unix())
	assert.Equal(t, now, user.CreatedAt())
	assert.Equal(t, now, user.UpdatedAt())
}

func TestRestoreUserModel(t *testing.T) {
	tests := []struct {
		name                   string
		id                     uint64
		userName               string
		email                  string
		password               string
		isActivated            bool
		confirmationToken      *string
		confirmationExpiresAt  *time.Time
		confirmedAt            *time.Time
		resetPasswordToken     *string
		resetPasswordExpiresAt *time.Time
		createdAt              time.Time
		updatedAt              time.Time
		expectError            bool
		errorMsg               string
	}{
		{
			name:                   "Valid user restoration",
			id:                     1,
			userName:               "John Doe",
			email:                  "test@example.com",
			password:               "hashedpassword",
			isActivated:            true,
			confirmationToken:      nil,
			confirmationExpiresAt:  nil,
			confirmedAt:            lo.ToPtr(time.Now().Add(-24 * time.Hour)),
			resetPasswordToken:     nil,
			resetPasswordExpiresAt: nil,
			createdAt:              time.Now().Add(-48 * time.Hour),
			updatedAt:              time.Now(),
			expectError:            false,
		},
		{
			name:                   "Valid user with confirmation data",
			id:                     2,
			userName:               "Jane Doe",
			email:                  "jane@example.com",
			password:               "hashedpassword",
			isActivated:            false,
			confirmationToken:      lo.ToPtr("token123"),
			confirmationExpiresAt:  lo.ToPtr(time.Now().Add(24 * time.Hour)),
			confirmedAt:            nil,
			resetPasswordToken:     nil,
			resetPasswordExpiresAt: nil,
			createdAt:              time.Now().Add(-1 * time.Hour),
			updatedAt:              time.Now(),
			expectError:            false,
		},
		{
			name:                   "Valid user with reset password data",
			id:                     3,
			userName:               "Bob Smith",
			email:                  "bob@example.com",
			password:               "hashedpassword",
			isActivated:            true,
			confirmationToken:      nil,
			confirmationExpiresAt:  nil,
			confirmedAt:            lo.ToPtr(time.Now().Add(-48 * time.Hour)),
			resetPasswordToken:     lo.ToPtr("reset123"),
			resetPasswordExpiresAt: lo.ToPtr(time.Now().Add(12 * time.Hour)),
			createdAt:              time.Now().Add(-72 * time.Hour),
			updatedAt:              time.Now(),
			expectError:            false,
		},
		{
			name:                   "Invalid email",
			id:                     4,
			userName:               "Invalid User",
			email:                  "invalid-email",
			password:               "hashedpassword",
			isActivated:            true,
			confirmationToken:      nil,
			confirmationExpiresAt:  nil,
			confirmedAt:            nil,
			resetPasswordToken:     nil,
			resetPasswordExpiresAt: nil,
			createdAt:              time.Now().Add(-24 * time.Hour),
			updatedAt:              time.Now(),
			expectError:            true,
			errorMsg:               "invalid email format",
		},
		{
			name:                   "Invalid name",
			id:                     5,
			userName:               "J", // Too short
			email:                  "test@example.com",
			password:               "hashedpassword",
			isActivated:            true,
			confirmationToken:      nil,
			confirmationExpiresAt:  nil,
			confirmedAt:            nil,
			resetPasswordToken:     nil,
			resetPasswordExpiresAt: nil,
			createdAt:              time.Now().Add(-24 * time.Hour),
			updatedAt:              time.Now(),
			expectError:            true,
			errorMsg:               "name must be at least 2 characters",
		},
		{
			name:                   "Empty password hash",
			id:                     6,
			userName:               "John Doe",
			email:                  "test@example.com",
			password:               "",
			isActivated:            true,
			confirmationToken:      nil,
			confirmationExpiresAt:  nil,
			confirmedAt:            nil,
			resetPasswordToken:     nil,
			resetPasswordExpiresAt: nil,
			createdAt:              time.Now().Add(-24 * time.Hour),
			updatedAt:              time.Now(),
			expectError:            true,
			errorMsg:               "password hash is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			user, err := RestoreUserModel(
				tc.id,
				tc.userName,
				tc.email,
				tc.password,
				tc.isActivated,
				tc.confirmationToken,
				tc.confirmationExpiresAt,
				tc.confirmedAt,
				tc.resetPasswordToken,
				tc.resetPasswordExpiresAt,
				tc.createdAt,
				tc.updatedAt,
			)

			if tc.expectError {
				assert.Error(t, err)
				if tc.errorMsg != "" {
					assert.Contains(t, err.Error(), tc.errorMsg)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tc.id, user.ID())
				assert.Equal(t, tc.email, user.Email())
				assert.Equal(t, tc.userName, user.Name())
				assert.Equal(t, tc.password, user.PasswordHash())
				assert.Equal(t, tc.isActivated, user.IsActivated())

				// Check confirmation data
				if tc.confirmationToken == nil {
					assert.Nil(t, user.ConfirmationToken())
				} else {
					assert.Equal(t, *tc.confirmationToken, *user.ConfirmationToken())
				}

				if tc.confirmationExpiresAt == nil {
					assert.Nil(t, user.ConfirmationExpiresAt())
				} else {
					assert.Equal(t, tc.confirmationExpiresAt.Unix(), user.ConfirmationExpiresAt().Unix())
				}

				if tc.confirmedAt == nil {
					assert.Nil(t, user.ConfirmedAt())
				} else {
					assert.Equal(t, tc.confirmedAt.Unix(), user.ConfirmedAt().Unix())
				}

				// Check reset password data
				if tc.resetPasswordToken == nil {
					assert.Nil(t, user.ResetPasswordToken())
				} else {
					assert.Equal(t, *tc.resetPasswordToken, *user.ResetPasswordToken())
				}

				if tc.resetPasswordExpiresAt == nil {
					assert.Nil(t, user.ResetPasswordExpiresAt())
				} else {
					assert.Equal(t, tc.resetPasswordExpiresAt.Unix(), user.ResetPasswordExpiresAt().Unix())
				}

				assert.Equal(t, tc.createdAt.Unix(), user.CreatedAt().Unix())
				assert.Equal(t, tc.updatedAt.Unix(), user.UpdatedAt().Unix())
			}
		})
	}
}

func TestConfirmationMethods(t *testing.T) {
	// Create a user with confirmation token and expiration time
	confirmationToken := "confirm_token"
	confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)
	user, err := CreateUserModel("Test User", "test@example.com", "password_hash", confirmationToken, confirmationExpiresAt)
	require.NoError(t, err)

	// Verify initial state
	assert.False(t, user.IsActivated())
	require.NotNil(t, user.ConfirmationToken())
	assert.Equal(t, confirmationToken, *user.ConfirmationToken())
	require.NotNil(t, user.ConfirmationExpiresAt())
	assert.Nil(t, user.ConfirmedAt())

	// Confirm the user
	user.ConfirmAccount()

	// Verify user is confirmed
	assert.True(t, user.IsActivated())
	assert.Nil(t, user.ConfirmationToken())
	assert.Nil(t, user.ConfirmationExpiresAt())
	require.NotNil(t, user.ConfirmedAt())

	// Confirming should update the updatedAt timestamp
	assert.True(t, user.UpdatedAt().After(user.CreatedAt()))
}

func TestResetPasswordMethods(t *testing.T) {
	// Create a user with confirmation token and expiration time
	confirmationToken := "confirm_token"
	confirmationExpiresAt := time.Now().UTC().Add(24 * time.Hour)
	user, err := CreateUserModel("Test User", "test@example.com", "password_hash", confirmationToken, confirmationExpiresAt)
	require.NoError(t, err)

	// Verify initial state
	assert.Nil(t, user.ResetPasswordToken())
	assert.Nil(t, user.ResetPasswordExpiresAt())

	// Set reset password details
	token := "reset_token"
	expiresAt := time.Now().UTC().Add(24 * time.Hour)
	user.SetResetPasswordDetails(token, expiresAt)

	// Verify reset password details were set
	require.NotNil(t, user.ResetPasswordToken())
	assert.Equal(t, token, *user.ResetPasswordToken())
	require.NotNil(t, user.ResetPasswordExpiresAt())
	assert.Equal(t, expiresAt.Unix(), user.ResetPasswordExpiresAt().Unix())

	// Clear reset password details
	user.ClearResetPasswordDetails()

	// Verify reset password details were cleared
	assert.Nil(t, user.ResetPasswordToken())
	assert.Nil(t, user.ResetPasswordExpiresAt())

	// Clearing should update the updatedAt timestamp
	assert.True(t, user.UpdatedAt().After(user.CreatedAt()))
}

func TestIsConfirmationTokenValid(t *testing.T) {
	// Create test cases for token validation
	tests := []struct {
		name           string
		token          string
		setupUser      func() UserModel
		expectedResult bool
	}{
		{
			name:  "Valid token",
			token: "valid_token",
			setupUser: func() UserModel {
				token := "valid_token"
				expiresAt := time.Now().UTC().Add(24 * time.Hour)
				user, _ := CreateUserModel("Test User", "test@example.com", "password_hash", token, expiresAt)
				return user
			},
			expectedResult: true,
		},
		{
			name:  "Invalid token",
			token: "invalid_token",
			setupUser: func() UserModel {
				token := "valid_token"
				expiresAt := time.Now().UTC().Add(24 * time.Hour)
				user, _ := CreateUserModel("Test User", "test@example.com", "password_hash", token, expiresAt)
				return user
			},
			expectedResult: false,
		},
		{
			name:  "Expired token",
			token: "valid_token",
			setupUser: func() UserModel {
				token := "valid_token"
				expiresAt := time.Now().UTC().Add(-24 * time.Hour) // Expired
				user, _ := CreateUserModel("Test User", "test@example.com", "password_hash", token, expiresAt)
				return user
			},
			expectedResult: false,
		},
		{
			name:  "Nil token",
			token: "valid_token",
			setupUser: func() UserModel {
				user, _ := CreateUserModel("Test User", "test@example.com", "password_hash", "token", time.Now().UTC().Add(24*time.Hour))
				user.confirmationToken = nil
				return user
			},
			expectedResult: false,
		},
		{
			name:  "Nil expiration",
			token: "valid_token",
			setupUser: func() UserModel {
				user, _ := CreateUserModel("Test User", "test@example.com", "password_hash", "valid_token", time.Now().UTC().Add(24*time.Hour))
				user.confirmationExpiresAt = nil
				return user
			},
			expectedResult: false,
		},
		{
			name:  "Already confirmed",
			token: "valid_token",
			setupUser: func() UserModel {
				user, _ := CreateUserModel("Test User", "test@example.com", "password_hash", "valid_token", time.Now().UTC().Add(24*time.Hour))
				user.ConfirmAccount() // This marks the account as confirmed
				return user
			},
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := tt.setupUser()
			result := user.IsConfirmationTokenValid(tt.token)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}
