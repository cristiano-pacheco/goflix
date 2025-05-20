package model

import (
	"errors"
	"strings"
	"unicode"
)

type EmailModel struct {
	value string
}

func CreateEmailModel(value string) (*EmailModel, error) {
	// Trim spaces before validation
	value = strings.TrimSpace(value)
	err := validateEmail(value)
	if err != nil {
		return nil, err
	}
	return &EmailModel{value: value}, nil
}

func (e *EmailModel) String() string {
	return e.value
}

func isValidLocalPartChar(c rune) bool {
	// Allow alphanumeric characters
	if unicode.IsLetter(c) || unicode.IsDigit(c) {
		return true
	}

	// Allow these special characters: !#$%&'*+-/=?^_`{|}~.
	switch c {
	case '!', '#', '$', '%', '&', '\'', '*', '+', '-', '/', '=', '?', '^', '_', '`', '{', '|', '}', '~', '.':
		return true
	}

	return false
}

func isValidDomainChar(c rune) bool {
	// Allow letters, digits, hyphens, and dots
	if unicode.IsLetter(c) || unicode.IsDigit(c) {
		return true
	}

	switch c {
	case '-', '.':
		return true
	}

	return false
}

func validateEmail(value string) error {
	// Value should already be trimmed in CreateEmailModel, but we'll trim again for safety
	value = strings.TrimSpace(value)
	if len(value) == 0 {
		return errors.New("email is required")
	}

	// Check for @ symbol
	atIndex := strings.Index(value, "@")
	if atIndex <= 0 || atIndex == len(value)-1 {
		return errors.New("invalid email format: missing @ symbol or invalid position")
	}

	// Get domain part for further checks
	domain := value[atIndex+1:]

	// Check for invalid characters in domain part first
	for _, char := range domain {
		if !isValidDomainChar(char) {
			return errors.New("invalid character in email domain part")
		}
	}

	// Check for dot in domain part
	dotIndex := strings.LastIndex(domain, ".")
	if dotIndex <= 0 || dotIndex == len(domain)-1 {
		return errors.New("invalid email format: missing dot in domain or invalid position")
	}

	// Check if domain starts with a dot
	if domain[0] == '.' {
		return errors.New("invalid email format: domain cannot start with a dot")
	}

	// Check local part (before @)
	localPart := value[:atIndex]
	if len(localPart) > 64 {
		return errors.New("local part of email exceeds maximum length of 64 characters")
	}

	// Check domain part
	if len(domain) > 255 {
		return errors.New("domain part of email exceeds maximum length of 255 characters")
	}

	// Check for invalid characters in local part
	for _, char := range localPart {
		if !isValidLocalPartChar(char) {
			return errors.New("invalid character in email local part")
		}
	}

	return nil
}
