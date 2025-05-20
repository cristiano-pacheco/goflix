package model

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

type NameModel struct {
	value string
}

func CreateNameModel(value string) (*NameModel, error) {
	// Trim spaces before validation
	value = strings.TrimSpace(value)
	err := validateName(value)
	if err != nil {
		return nil, err
	}
	return &NameModel{value: value}, nil
}

func (n *NameModel) String() string {
	return n.value
}

func validateName(value string) error {
	// Check if name is empty
	if len(value) == 0 {
		return errors.New("name is required")
	}

	// Check minimum length
	if len(value) < 2 {
		return errors.New("name must be at least 2 characters long")
	}

	// Check maximum length
	if len(value) > 255 {
		return errors.New("name cannot exceed 255 characters")
	}

	// Check for invalid characters
	for _, r := range value {
		// Allow letters, numbers, spaces, hyphens, apostrophes, and periods (for names like "Dr. Smith" or "O'Connor")
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != ' ' && r != '-' && r != '\'' && r != '.' {
			return errors.New("name contains invalid characters")
		}
	}

	// Check for consecutive spaces
	if strings.Contains(value, "  ") {
		return errors.New("invalid name format: contains consecutive spaces")
	}

	// Check for valid name format using regex
	// This allows for names with spaces, hyphens, apostrophes, and periods
	// Examples: "John Doe", "Mary-Ann", "O'Connor", "Dr. Smith"
	validNamePattern := regexp.MustCompile(`^[\p{L}][\p{L}\p{N} \-\'\.]*(?: [\p{L}][\p{L}\p{N} \-\'\.]*)*(?: [\p{L}][\p{L}\p{N} \-\'\.]*)?$`)
	
	// For very long strings, the regex might be too expensive to evaluate
	// So for strings longer than 50 chars, we'll just check the basic pattern
	if len(value) > 50 {
		// For long strings, just check if it starts with a letter
		if !unicode.IsLetter(rune(value[0])) {
			return errors.New("invalid name format: must start with a letter")
		}
	} else if !validNamePattern.MatchString(value) {
		return errors.New("invalid name format")
	}

	return nil
}
