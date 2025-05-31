package model

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/samber/lo"
)

const (
	minNameModelLength            = 2
	maxNameModelLength            = 100
	maxNameConsecutivePunctuation = 2
)

type NameModel struct {
	value string
}

func CreateNameModel(value string) (NameModel, error) {
	value = strings.TrimSpace(value)
	if err := validatePlanName(value); err != nil {
		return NameModel{}, err
	}
	return NameModel{value: value}, nil
}

func (n NameModel) String() string {
	return n.value
}

func validatePlanName(value string) error {
	charCount := utf8.RuneCountInString(value)

	if charCount == 0 {
		return errors.New("plan name is required")
	}

	if charCount < minNameModelLength {
		return errors.New("plan name must be at least 2 characters long")
	}

	if charCount > maxNameModelLength {
		return errors.New("plan name cannot exceed 100 characters")
	}

	// Check if name starts with a letter or digit
	firstRune := []rune(value)[0]
	if !unicode.IsLetter(firstRune) && !unicode.IsDigit(firstRune) {
		return errors.New("plan name must start with a letter or digit")
	}

	// Check if name ends with a letter or digit (not punctuation)
	lastRune := []rune(value)[len([]rune(value))-1]
	if !unicode.IsLetter(lastRune) && !unicode.IsDigit(lastRune) {
		return errors.New("plan name must end with a letter or digit")
	}

	// Check for consecutive spaces
	if strings.Contains(value, "  ") {
		return errors.New("plan name cannot contain consecutive spaces")
	}

	// Check for invalid characters using functional approach
	runes := []rune(value)
	if !lo.EveryBy(runes, isValidPlanNameChar) {
		return errors.New(
			"plan name contains invalid characters (only letters, digits, spaces, hyphens, underscores, and periods are allowed)",
		)
	}

	// Additional format validations
	if err := validatePlanNameFormat(value); err != nil {
		return err
	}

	return nil
}

func validatePlanNameFormat(value string) error {
	if err := validatePlanNameBoundaries(value); err != nil {
		return err
	}

	return validatePlanNamePunctuationRules(value)
}

func validatePlanNameBoundaries(value string) error {
	// Check for leading or trailing spaces (should be trimmed already, but double-check)
	if strings.HasPrefix(value, " ") || strings.HasSuffix(value, " ") {
		return errors.New("plan name cannot start or end with spaces")
	}

	// Check for leading or trailing punctuation
	if strings.HasPrefix(value, "-") || strings.HasPrefix(value, "_") || strings.HasPrefix(value, ".") {
		return errors.New("plan name cannot start with punctuation")
	}

	if strings.HasSuffix(value, "-") || strings.HasSuffix(value, "_") || strings.HasSuffix(value, ".") {
		return errors.New("plan name cannot end with punctuation")
	}

	return nil
}

func validatePlanNamePunctuationRules(value string) error {
	// Check for excessive punctuation (more than 2 consecutive punctuation marks)
	punctuationCount := 0
	for _, r := range value {
		if r == '-' || r == '_' || r == '.' {
			punctuationCount++
			if punctuationCount > maxNameConsecutivePunctuation {
				return errors.New("plan name cannot contain more than 2 consecutive punctuation marks")
			}
		} else {
			punctuationCount = 0
		}
	}

	return nil
}

func isValidPlanNameChar(r rune) bool {
	// Allow letters, digits, spaces, hyphens, underscores, and periods
	return unicode.IsLetter(r) ||
		unicode.IsDigit(r) ||
		r == ' ' ||
		r == '-' ||
		r == '_' ||
		r == '.'
}
