package model

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/samber/lo"
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

func (d DescriptionModel) String() string {
	return d.value
}

func validateDescription(value string) error {
	// Description is optional, so empty is allowed
	if value == "" {
		return nil
	}

	charCount := utf8.RuneCountInString(value)

	if charCount > maxDescriptionLength {
		return errors.New("description cannot exceed 255 characters")
	}

	// Check for invalid characters using functional approach
	runes := []rune(value)
	if !lo.EveryBy(runes, isValidDescriptionChar) {
		return errors.New(
			"description contains invalid characters (only printable characters are allowed)",
		)
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
		return errors.New("description cannot start or end with spaces")
	}

	// Check for excessive consecutive spaces
	if strings.Contains(value, "   ") {
		return errors.New("description cannot contain more than 2 consecutive spaces")
	}

	// Check for control characters (except newlines and tabs which might be useful)
	for _, r := range value {
		if unicode.IsControl(r) && r != '\n' && r != '\t' {
			return errors.New("description cannot contain control characters")
		}
	}

	return nil
}

func isValidDescriptionChar(r rune) bool {
	// Allow all printable characters, newlines, and tabs
	return unicode.IsPrint(r) || r == '\n' || r == '\t'
}
