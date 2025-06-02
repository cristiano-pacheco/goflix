package model

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/samber/lo"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
)

const (
	maxDescriptionLength = 1000
)

type DescriptionModel struct {
	value string
}

func CreateDescriptionModel(value string) (DescriptionModel, error) {
	value = strings.TrimSpace(value)
	if err := validateDescription(value); err != nil {
		return DescriptionModel{}, err
	}
	return DescriptionModel{value: value}, nil
}

func (d *DescriptionModel) String() string {
	return d.value
}

func validateDescription(value string) error {
	// Description is optional, so empty is allowed
	if value == "" {
		return nil
	}

	charCount := utf8.RuneCountInString(value)

	if charCount > maxDescriptionLength {
		return errs.ErrDescriptionTooLong
	}

	// Check for invalid characters using functional approach
	runes := []rune(value)
	if !lo.EveryBy(runes, isValidDescriptionChar) {
		return errs.ErrDescriptionInvalidCharacters
	}

	// Additional format validations
	if err := validateDescriptionFormat(value); err != nil {
		return err
	}

	return nil
}

func validateDescriptionFormat(value string) error {
	// Check for leading or trailing spaces (should be trimmed already, but double-check)
	if strings.HasPrefix(value, " ") || strings.HasSuffix(value, " ") {
		return errs.ErrDescriptionCannotStartOrEndWithSpaces
	}

	// Check for excessive consecutive spaces
	if strings.Contains(value, "   ") {
		return errs.ErrDescriptionExcessiveConsecutiveSpaces
	}

	// Check for control characters (except newlines and tabs which might be useful)
	for _, r := range value {
		if unicode.IsControl(r) && r != '\n' && r != '\t' {
			return errs.ErrDescriptionControlCharacters
		}
	}

	return nil
}

func isValidDescriptionChar(r rune) bool {
	// Allow all printable characters, newlines, and tabs
	return unicode.IsPrint(r) || r == '\n' || r == '\t'
}
