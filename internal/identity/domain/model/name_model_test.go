package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateNameModel_ValidNames(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple first name",
			input:    "John",
			expected: "John",
		},
		{
			name:     "full name with space",
			input:    "John Doe",
			expected: "John Doe",
		},
		{
			name:     "name with hyphen",
			input:    "Mary-Ann",
			expected: "Mary-Ann",
		},
		{
			name:     "name with apostrophe",
			input:    "O'Connor",
			expected: "O'Connor",
		},
		{
			name:     "name with title and period",
			input:    "Dr. Smith",
			expected: "Dr. Smith",
		},
		{
			name:     "name with multiple parts",
			input:    "Jean-Claude Van Damme",
			expected: "Jean-Claude Van Damme",
		},
		{
			name:     "name with numbers",
			input:    "John Doe Jr2",
			expected: "John Doe Jr2",
		},
		{
			name:     "name with leading/trailing spaces (should be trimmed)",
			input:    "  John Doe  ",
			expected: "John Doe",
		},
		{
			name:     "minimum length name",
			input:    "Jo",
			expected: "Jo",
		},
		{
			name:     "name with period in middle",
			input:    "J.R.R. Tolkien",
			expected: "J.R.R. Tolkien",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nameModel, err := CreateNameModel(tc.input)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, nameModel.String())
		})
	}
}

func TestCreateNameModel_InvalidNames(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError string
	}{
		{
			name:          "empty name",
			input:         "",
			expectedError: "name is required",
		},
		{
			name:          "only spaces",
			input:         "   ",
			expectedError: "name is required",
		},
		{
			name:          "too short",
			input:         "J",
			expectedError: "name must be at least 2 characters long",
		},
		{
			name:          "too long",
			input:         strings.Repeat("a", 256),
			expectedError: "name cannot exceed 255 characters",
		},
		{
			name:          "starts with number",
			input:         "1John",
			expectedError: "name must start with a letter",
		},
		{
			name:          "starts with space",
			input:         " John",
			expectedError: "name must start with a letter",
		},
		{
			name:          "starts with hyphen",
			input:         "-John",
			expectedError: "name cannot start with punctuation",
		},
		{
			name:          "starts with apostrophe",
			input:         "'John",
			expectedError: "name cannot start with punctuation",
		},
		{
			name:          "ends with hyphen",
			input:         "John-",
			expectedError: "name cannot end with punctuation",
		},
		{
			name:          "ends with apostrophe",
			input:         "John'",
			expectedError: "name cannot end with punctuation",
		},
		{
			name:          "ends with period",
			input:         "John.",
			expectedError: "name must end with a letter or digit",
		},
		{
			name:          "consecutive spaces",
			input:         "John  Doe",
			expectedError: "name cannot contain consecutive spaces",
		},
		{
			name:          "invalid characters - special symbols",
			input:         "John@Doe",
			expectedError: "name contains invalid characters (only letters, digits, spaces, hyphens, apostrophes, and periods are allowed)",
		},
		{
			name:          "invalid characters - underscore",
			input:         "John_Doe",
			expectedError: "name contains invalid characters (only letters, digits, spaces, hyphens, apostrophes, and periods are allowed)",
		},
		{
			name:          "invalid characters - parentheses",
			input:         "John (Doe)",
			expectedError: "name contains invalid characters (only letters, digits, spaces, hyphens, apostrophes, and periods are allowed)",
		},
		{
			name:          "excessive punctuation",
			input:         "John----Doe",
			expectedError: "name cannot contain more than 3 consecutive punctuation marks",
		},
		{
			name:          "excessive apostrophes",
			input:         "John''''Doe",
			expectedError: "name cannot contain more than 3 consecutive punctuation marks",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nameModel, err := CreateNameModel(tc.input)
			require.Error(t, err)
			assert.Equal(t, NameModel{}, nameModel)
			assert.Contains(t, err.Error(), tc.expectedError)
		})
	}
}

func TestNameModel_String(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple name",
			input:    "John",
			expected: "John",
		},
		{
			name:     "full name",
			input:    "John Doe",
			expected: "John Doe",
		},
		{
			name:     "complex name",
			input:    "Jean-Claude O'Connor Jr.",
			expected: "Jean-Claude O'Connor Jr.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nameModel, err := CreateNameModel(tc.input)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, nameModel.String())
		})
	}
}

func TestValidateNameFormat(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expectedError string
	}{
		{
			name:          "valid name",
			input:         "John Doe",
			expectedError: "",
		},
		{
			name:          "leading space",
			input:         " John",
			expectedError: "name cannot start or end with spaces",
		},
		{
			name:          "trailing space",
			input:         "John ",
			expectedError: "name cannot start or end with spaces",
		},
		{
			name:          "starts with hyphen",
			input:         "-John",
			expectedError: "name cannot start with punctuation",
		},
		{
			name:          "ends with hyphen",
			input:         "John-",
			expectedError: "name cannot end with punctuation",
		},
		{
			name:          "excessive punctuation",
			input:         "John----Doe",
			expectedError: "name cannot contain more than 3 consecutive punctuation marks",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateNameFormat(tc.input)
			if tc.expectedError == "" {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			}
		})
	}
}

func TestIsValidNameChar(t *testing.T) {
	testCases := []struct {
		name     string
		input    rune
		expected bool
	}{
		{
			name:     "letter",
			input:    'A',
			expected: true,
		},
		{
			name:     "lowercase letter",
			input:    'a',
			expected: true,
		},
		{
			name:     "digit",
			input:    '1',
			expected: true,
		},
		{
			name:     "space",
			input:    ' ',
			expected: true,
		},
		{
			name:     "hyphen",
			input:    '-',
			expected: true,
		},
		{
			name:     "apostrophe",
			input:    '\'',
			expected: true,
		},
		{
			name:     "period",
			input:    '.',
			expected: true,
		},
		{
			name:     "at symbol",
			input:    '@',
			expected: false,
		},
		{
			name:     "underscore",
			input:    '_',
			expected: false,
		},
		{
			name:     "parenthesis",
			input:    '(',
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidNameChar(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// Benchmark tests
func BenchmarkCreateNameModel(b *testing.B) {
	testName := "John Doe"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = CreateNameModel(testName)
	}
}

func BenchmarkValidateName(b *testing.B) {
	testName := "John Doe"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = validateName(testName)
	}
}
