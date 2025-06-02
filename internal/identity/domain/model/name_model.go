package model

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/samber/lo"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/errs"
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
		return errs.ErrNameRequired
	}

	if charCount < minNameLength {
		return errs.ErrNameTooShort
	}

	if charCount > maxNameLength {
		return errs.ErrNameTooLong
	}

	// Check if name starts with a letter
	firstRune := []rune(value)[0]
	if !unicode.IsLetter(firstRune) {
		return errs.ErrNameMustStartWithLetter
	}

	// Check if name ends with a letter or digit (not punctuation)
	lastRune := []rune(value)[len([]rune(value))-1]
	if !unicode.IsLetter(lastRune) && !unicode.IsDigit(lastRune) {
		return errs.ErrNameMustEndWithLetterOrDigit
	}

	// Check for consecutive spaces
	if strings.Contains(value, "  ") {
		return errs.ErrNameConsecutiveSpaces
	}

	// Check for invalid characters using functional approach
	runes := []rune(value)
	if !lo.EveryBy(runes, isValidNameChar) {
		return errs.ErrNameInvalidCharacters
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
		return errs.ErrNameCannotStartOrEndWithSpaces
	}

	// Check for leading or trailing punctuation (except for titles like "Dr.")
	if strings.HasPrefix(value, "-") || strings.HasPrefix(value, "'") {
		return errs.ErrNameCannotStartWithPunctuation
	}

	if strings.HasSuffix(value, "-") || strings.HasSuffix(value, "'") {
		return errs.ErrNameCannotEndWithPunctuation
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
				return errs.ErrNameExcessiveConsecutivePunctuation
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
