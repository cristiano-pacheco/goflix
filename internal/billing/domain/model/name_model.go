package model

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/samber/lo"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
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

func (n *NameModel) String() string {
	return n.value
}

func validatePlanName(value string) error {
	if err := validatePlanNameLength(value); err != nil {
		return err
	}

	if err := validatePlanNameCharacters(value); err != nil {
		return err
	}

	if err := validatePlanNameFormat(value); err != nil {
		return err
	}

	return nil
}

func validatePlanNameLength(value string) error {
	charCount := utf8.RuneCountInString(value)

	if charCount == 0 {
		return errs.ErrNameRequired
	}

	if charCount < minNameModelLength {
		return errs.ErrNameTooShort
	}

	if charCount > maxNameModelLength {
		return errs.ErrNameTooLong
	}

	return nil
}

func validatePlanNameCharacters(value string) error {
	runes := []rune(value)

	// Check if name starts with a letter or digit
	if !unicode.IsLetter(runes[0]) && !unicode.IsDigit(runes[0]) {
		return errs.ErrNameMustStartWithLetterOrDigit
	}

	// Check if name ends with a letter or digit (not punctuation)
	lastRune := runes[len(runes)-1]
	if !unicode.IsLetter(lastRune) && !unicode.IsDigit(lastRune) {
		return errs.ErrNameMustEndWithLetterOrDigit
	}

	// Check for consecutive spaces
	if strings.Contains(value, "  ") {
		return errs.ErrNameConsecutiveSpaces
	}

	// Check for invalid characters using functional approach
	if !lo.EveryBy(runes, isValidPlanNameChar) {
		return errs.ErrNameInvalidCharacters
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
		return errs.ErrNameCannotStartOrEndWithSpaces
	}

	// Check for leading or trailing punctuation
	if strings.HasPrefix(value, "-") || strings.HasPrefix(value, "_") || strings.HasPrefix(value, ".") {
		return errs.ErrNameCannotStartWithPunctuation
	}

	if strings.HasSuffix(value, "-") || strings.HasSuffix(value, "_") || strings.HasSuffix(value, ".") {
		return errs.ErrNameCannotEndWithPunctuation
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
				return errs.ErrNameExcessiveConsecutivePunctuation
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
