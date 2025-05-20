package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateEmailModel(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid email",
			email:       "test@example.com",
			expectError: false,
		},
		{
			name:        "Empty email",
			email:       "",
			expectError: true,
			errorMsg:    "email is required",
		},
		{
			name:        "Email with spaces",
			email:       "  test@example.com  ",
			expectError: false,
		},
		{
			name:        "Email without @",
			email:       "testexample.com",
			expectError: true,
			errorMsg:    "invalid email format: missing @ symbol or invalid position",
		},
		{
			name:        "Email with @ at beginning",
			email:       "@example.com",
			expectError: true,
			errorMsg:    "invalid email format: missing @ symbol or invalid position",
		},
		{
			name:        "Email with @ at end",
			email:       "test@",
			expectError: true,
			errorMsg:    "invalid email format: missing @ symbol or invalid position",
		},
		{
			name:        "Email without dot in domain",
			email:       "test@examplecom",
			expectError: true,
			errorMsg:    "invalid email format: missing dot in domain or invalid position",
		},
		{
			name:        "Email with dot at beginning of domain",
			email:       "test@.example.com",
			expectError: true,
			errorMsg:    "invalid email format: domain cannot start with a dot",
		},
		{
			name:        "Email with dot at end of domain",
			email:       "test@example.",
			expectError: true,
			errorMsg:    "invalid email format: missing dot in domain or invalid position",
		},
		{
			name:        "Email with very long local part",
			email:       strings.Repeat("a", 65) + "@example.com",
			expectError: true,
			errorMsg:    "local part of email exceeds maximum length of 64 characters",
		},
		{
			name:        "Email with very long domain part",
			email:       "test@" + strings.Repeat("a", 256) + ".com",
			expectError: true,
			errorMsg:    "domain part of email exceeds maximum length of 255 characters",
		},
		{
			name:        "Email with invalid character in local part",
			email:       "test<script>@example.com",
			expectError: true,
			errorMsg:    "invalid character in email local part",
		},
		{
			name:        "Email with invalid character in domain part",
			email:       "test@example_com",
			expectError: true,
			errorMsg:    "invalid character in email domain part",
		},
		{
			name:        "Email with special characters in local part",
			email:       "test.user+filter!#$%&'*-/=?^_`{|}~@example.com",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := CreateEmailModel(tt.email)

			if tt.expectError {
				assert.Error(t, err, "Expected an error for invalid email: %s", tt.email)
				assert.Nil(t, email, "Expected nil email model for invalid email: %s", tt.email)
				if err != nil { // Avoid nil pointer dereference
					assert.Contains(t, err.Error(), tt.errorMsg, "Error message doesn't match for email: %s", tt.email)
				}
			} else {
				assert.NoError(t, err, "Unexpected error for valid email: %s", tt.email)
				require.NotNil(t, email, "Email model should not be nil for valid email: %s", tt.email)
				if email != nil { // Avoid nil pointer dereference
					// If the input had spaces, they should be trimmed
					expectedValue := strings.TrimSpace(tt.email)
					assert.Equal(t, expectedValue, email.String(), "Email value doesn't match for: %s", tt.email)
				}
			}
		})
	}
}

func TestEmailModel_String(t *testing.T) {
	// Create a valid email model
	email, err := CreateEmailModel("test@example.com")
	require.NoError(t, err)
	require.NotNil(t, email)

	// Test the String method
	assert.Equal(t, "test@example.com", email.String())
}
