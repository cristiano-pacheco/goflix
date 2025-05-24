package model

import (
	"errors"
	"strings"
	"unicode"

	"github.com/samber/lo"
)

type EmailModel struct {
	value string
}

func CreateEmailModel(value string) (EmailModel, error) {
	value = strings.TrimSpace(value)
	if err := validateEmail(value); err != nil {
		return EmailModel{}, err
	}
	return EmailModel{value: value}, nil
}

func (e EmailModel) String() string {
	return e.value
}

func validateEmail(value string) error {
	if len(value) == 0 {
		return errors.New("email is required")
	}

	// Check overall length to prevent potential DoS attacks
	if len(value) > 320 { // RFC 5321 limit
		return errors.New("email exceeds maximum length of 320 characters")
	}

	// Find @ symbol
	atIndex := strings.Index(value, "@")
	if atIndex <= 0 || atIndex == len(value)-1 {
		return errors.New("invalid email format: missing @ symbol or invalid position")
	}

	// Check for multiple @ symbols
	if strings.Count(value, "@") != 1 {
		return errors.New("invalid email format: multiple @ symbols found")
	}

	localPart := value[:atIndex]
	domain := value[atIndex+1:]

	// Validate local part
	if err := validateLocalPart(localPart); err != nil {
		return err
	}

	// Validate domain part
	if err := validateDomain(domain); err != nil {
		return err
	}

	return nil
}

func validateLocalPart(localPart string) error {
	if len(localPart) == 0 {
		return errors.New("email local part cannot be empty")
	}

	if len(localPart) > 64 {
		return errors.New("email local part exceeds maximum length of 64 characters")
	}

	// Check for consecutive dots
	if strings.Contains(localPart, "..") {
		return errors.New("email local part cannot contain consecutive dots")
	}

	// Check if starts or ends with dot
	if localPart[0] == '.' || localPart[len(localPart)-1] == '.' {
		return errors.New("email local part cannot start or end with a dot")
	}

	// Validate characters
	for _, char := range localPart {
		if !isValidLocalPartChar(char) {
			return errors.New("email local part contains invalid characters")
		}
	}

	return nil
}

func validateDomain(domain string) error {
	if len(domain) == 0 {
		return errors.New("email domain cannot be empty")
	}

	if len(domain) > 255 {
		return errors.New("email domain exceeds maximum length of 255 characters")
	}

	// Check if starts or ends with dot
	if domain[0] == '.' || domain[len(domain)-1] == '.' {
		return errors.New("email domain cannot start or end with a dot")
	}

	// Check if starts or ends with hyphen
	if domain[0] == '-' || domain[len(domain)-1] == '-' {
		return errors.New("email domain cannot start or end with a hyphen")
	}

	// Must contain at least one dot
	if !strings.Contains(domain, ".") {
		return errors.New("email domain must contain at least one dot")
	}

	// Check for consecutive dots
	if strings.Contains(domain, "..") {
		return errors.New("email domain cannot contain consecutive dots")
	}

	// Split domain into labels and validate each
	labels := strings.Split(domain, ".")
	if len(labels) < 2 {
		return errors.New("email domain must have at least two labels")
	}

	// Validate each label
	for i, label := range labels {
		if err := validateDomainLabel(label, i == len(labels)-1); err != nil {
			return err
		}
	}

	return nil
}

func validateDomainLabel(label string, isTopLevel bool) error {
	if len(label) == 0 {
		return errors.New("email domain label cannot be empty")
	}

	if len(label) > 63 {
		return errors.New("email domain label exceeds maximum length of 63 characters")
	}

	// Top-level domain should be at least 2 characters and contain only letters
	if isTopLevel {
		if len(label) < 2 {
			return errors.New("email top-level domain must be at least 2 characters")
		}

		// Check if all characters are letters (more restrictive for TLD)
		if !lo.EveryBy([]rune(label), func(char rune) bool {
			return unicode.IsLetter(char)
		}) {
			return errors.New("email top-level domain must contain only letters")
		}
	} else {
		// Regular domain labels can contain letters, digits, and hyphens
		// but cannot start or end with hyphen
		if label[0] == '-' || label[len(label)-1] == '-' {
			return errors.New("email domain label cannot start or end with hyphen")
		}

		if !lo.EveryBy([]rune(label), isValidDomainChar) {
			return errors.New("email domain label contains invalid characters")
		}
	}

	return nil
}

func isValidLocalPartChar(c rune) bool {
	// Allow alphanumeric characters
	if unicode.IsLetter(c) || unicode.IsDigit(c) {
		return true
	}

	// Allow RFC 5322 special characters for local part
	allowedSpecialChars := []rune{'!', '#', '$', '%', '&', '\'', '*', '+', '-', '/', '=', '?', '^', '_', '`', '{', '|', '}', '~', '.'}
	return lo.Contains(allowedSpecialChars, c)
}

func isValidDomainChar(c rune) bool {
	// Allow letters, digits, and hyphens for domain labels
	return unicode.IsLetter(c) || unicode.IsDigit(c) || c == '-'
}
