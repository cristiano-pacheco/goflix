package model

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/samber/lo"
)

const (
	minNameLength             = 2
	maxNameLength             = 255
	maxConsecutivePunctuation = 3
)

type NameModel struct {
	value string
}

func CreateNameModel(value string) (NameModel, error) {
	value = strings.TrimSpace(value)
	if err := validateName(value); err != nil {
		return NameModel{}, err
	}
	return NameModel{value: value}, nil
}

func (n NameModel) String() string {
	return n.value
}

func validateName(value string) error {
	charCount := utf8.RuneCountInString(value)

	if charCount == 0 {
		return errors.New("name is required")
	}

	if charCount < minNameLength {
		return errors.New("name must be at least 2 characters long")
	}

	if charCount > maxNameLength {
		return errors.New("name cannot exceed 255 characters")
	}

	// Check if name starts with a letter
	firstRune := []rune(value)[0]
	if !unicode.IsLetter(firstRune) {
		return errors.New("name must start with a letter")
	}

	// Check if name ends with a letter or digit (not punctuation)
	lastRune := []rune(value)[len([]rune(value))-1]
	if !unicode.IsLetter(lastRune) && !unicode.IsDigit(lastRune) {
		return errors.New("name must end with a letter or digit")
	}

	// Check for consecutive spaces
	if strings.Contains(value, "  ") {
		return errors.New("name cannot contain consecutive spaces")
	}

	// Check for invalid characters using functional approach
	runes := []rune(value)
	if !lo.EveryBy(runes, isValidNameChar) {
		return errors.New(
			"name contains invalid characters (only letters, digits, spaces, hyphens, apostrophes, and periods are allowed)",
		)
	}

	// Additional format validations
	if err := validateNameFormat(value); err != nil {
		return err
	}

	return nil
}

func validateNameFormat(value string) error {
	if err := validateNameBoundaries(value); err != nil {
		return err
	}

	return validatePunctuationRules(value)
}

func validateNameBoundaries(value string) error {
	// Check for leading or trailing spaces (should be trimmed already, but double-check)
	if strings.HasPrefix(value, " ") || strings.HasSuffix(value, " ") {
		return errors.New("name cannot start or end with spaces")
	}

	// Check for leading or trailing punctuation (except for titles like "Dr.")
	if strings.HasPrefix(value, "-") || strings.HasPrefix(value, "'") {
		return errors.New("name cannot start with punctuation")
	}

	if strings.HasSuffix(value, "-") || strings.HasSuffix(value, "'") {
		return errors.New("name cannot end with punctuation")
	}

	return nil
}

func validatePunctuationRules(value string) error {
	// Check for excessive punctuation (more than 3 consecutive punctuation marks)
	punctuationCount := 0
	for _, r := range value {
		if r == '-' || r == '\'' || r == '.' {
			punctuationCount++
			if punctuationCount > maxConsecutivePunctuation {
				return errors.New("name cannot contain more than 3 consecutive punctuation marks")
			}
		} else {
			punctuationCount = 0
		}
	}

	return nil
}

func isValidNameChar(r rune) bool {
	// Allow letters, digits, spaces, hyphens, apostrophes, and periods
	return unicode.IsLetter(r) ||
		unicode.IsDigit(r) ||
		r == ' ' ||
		r == '-' ||
		r == '\'' ||
		r == '.'
}
