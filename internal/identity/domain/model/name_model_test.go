package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateNameModel(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid name",
			input:       "John Doe",
			expectError: false,
		},
		{
			name:        "Empty name",
			input:       "",
			expectError: true,
			errorMsg:    "name is required",
		},
		{
			name:        "Name with spaces",
			input:       "  Jane Smith  ",
			expectError: false,
		},
		{
			name:        "Name too short",
			input:       "J",
			expectError: true,
			errorMsg:    "name must be at least 2 characters long",
		},
		{
			name:        "Name at minimum length",
			input:       "Jo",
			expectError: false,
		},
		{
			name:        "Name at maximum valid length",
			input:       "John " + strings.Repeat("a", 250),
			expectError: false,
		},
		{
			name:        "Name with invalid format (long string)",
			input:       "1" + strings.Repeat("a", 254), // Starts with a number
			expectError: true,
			errorMsg:    "invalid name format: must start with a letter",
		},
		{
			name:        "Name too long",
			input:       strings.Repeat("a", 256),
			expectError: true,
			errorMsg:    "name cannot exceed 255 characters",
		},
		{
			name:        "Name with hyphen",
			input:       "Mary-Jane Watson",
			expectError: false,
		},
		{
			name:        "Name with apostrophe",
			input:       "O'Connor",
			expectError: false,
		},
		{
			name:        "Name with period",
			input:       "Dr. John Smith",
			expectError: false,
		},
		{
			name:        "Name with invalid character",
			input:       "John@Doe",
			expectError: true,
			errorMsg:    "name contains invalid characters",
		},
		{
			name:        "Name starting with number",
			input:       "1John",
			expectError: true,
			errorMsg:    "invalid name format",
		},
		{
			name:        "Name with valid number in middle",
			input:       "John2Doe",
			expectError: false,
		},
		{
			name:        "Name with invalid format",
			input:       "John  Doe", // Double space
			expectError: true,
			errorMsg:    "invalid name format: contains consecutive spaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nameModel, err := CreateNameModel(tt.input)

			if tt.expectError {
				assert.Error(t, err, "Expected an error for invalid name: %s", tt.input)
				assert.Nil(t, nameModel, "Expected nil name model for invalid name: %s", tt.input)
				if err != nil { // Avoid nil pointer dereference
					assert.Contains(t, err.Error(), tt.errorMsg, "Error message doesn't match for name: %s", tt.input)
				}
			} else {
				assert.NoError(t, err, "Unexpected error for valid name: %s", tt.input)
				require.NotNil(t, nameModel, "Name model should not be nil for valid name: %s", tt.input)
				if nameModel != nil { // Avoid nil pointer dereference
					// If the input had spaces, they should be trimmed
					expectedValue := strings.TrimSpace(tt.input)
					assert.Equal(t, expectedValue, nameModel.String(), "Name value doesn't match for: %s", tt.input)
				}
			}
		})
	}
}

func TestNameModel_String(t *testing.T) {
	// Create a valid name model
	nameModel, err := CreateNameModel("John Doe")
	require.NoError(t, err)
	require.NotNil(t, nameModel)

	// Test the String method
	assert.Equal(t, "John Doe", nameModel.String())
}
