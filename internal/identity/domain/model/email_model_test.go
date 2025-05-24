package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmailModel(t *testing.T) {
	t.Run("valid email addresses", func(t *testing.T) {
		validEmails := []string{
			"test@example.com",
			"user.name@domain.co.uk",
			"user+tag@example.org",
			"user_name@example-domain.com",
			"a@b.co",
			"test123@example123.com",
			"user.name+tag@sub.domain.com",
		}

		for _, email := range validEmails {
			// Arrange & Act
			result, err := CreateEmailModel(email)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, email, result.String())
		}
	})

	t.Run("valid email with whitespace trimming", func(t *testing.T) {
		// Arrange
		emailWithSpaces := "  test@example.com  "
		expectedEmail := "test@example.com"

		// Act
		result, err := CreateEmailModel(emailWithSpaces)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedEmail, result.String())
	})

	t.Run("empty email", func(t *testing.T) {
		// Arrange
		email := ""

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email is required", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("whitespace only email", func(t *testing.T) {
		// Arrange
		email := "   "

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email is required", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("email exceeds maximum length", func(t *testing.T) {
		// Arrange
		longEmail := "a" + "@" + "b" + ".com"
		for len(longEmail) <= 320 {
			longEmail = "a" + longEmail
		}

		// Act
		result, err := CreateEmailModel(longEmail)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email exceeds maximum length of 320 characters", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("missing @ symbol", func(t *testing.T) {
		// Arrange
		email := "testexample.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "invalid email format: missing @ symbol or invalid position", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("@ symbol at beginning", func(t *testing.T) {
		// Arrange
		email := "@example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "invalid email format: missing @ symbol or invalid position", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("@ symbol at end", func(t *testing.T) {
		// Arrange
		email := "test@"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "invalid email format: missing @ symbol or invalid position", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("multiple @ symbols", func(t *testing.T) {
		// Arrange
		email := "test@example@com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "invalid email format: multiple @ symbols found", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("local part too long", func(t *testing.T) {
		// Arrange
		longLocalPart := ""
		for len(longLocalPart) <= 64 {
			longLocalPart += "a"
		}
		email := longLocalPart + "@example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email local part exceeds maximum length of 64 characters", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("local part with consecutive dots", func(t *testing.T) {
		// Arrange
		email := "test..user@example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email local part cannot contain consecutive dots", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("local part starts with dot", func(t *testing.T) {
		// Arrange
		email := ".test@example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email local part cannot start or end with a dot", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("local part ends with dot", func(t *testing.T) {
		// Arrange
		email := "test.@example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email local part cannot start or end with a dot", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("local part with invalid characters", func(t *testing.T) {
		// Arrange
		email := "test@user@example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "invalid email format: multiple @ symbols found", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain too long", func(t *testing.T) {
		// Arrange
		longDomain := ""
		for len(longDomain) <= 255 {
			longDomain += "a"
		}
		longDomain += ".com"
		email := "test@" + longDomain

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain exceeds maximum length of 255 characters", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain starts with dot", func(t *testing.T) {
		// Arrange
		email := "test@.example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain cannot start or end with a dot", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain ends with dot", func(t *testing.T) {
		// Arrange
		email := "test@example.com."

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain cannot start or end with a dot", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain starts with hyphen", func(t *testing.T) {
		// Arrange
		email := "test@-example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain cannot start or end with a hyphen", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain ends with hyphen", func(t *testing.T) {
		// Arrange
		email := "test@example-.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain cannot start or end with a hyphen", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain without dot", func(t *testing.T) {
		// Arrange
		email := "test@example"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain must contain at least one dot", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain with consecutive dots", func(t *testing.T) {
		// Arrange
		email := "test@example..com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain cannot contain consecutive dots", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain label too long", func(t *testing.T) {
		// Arrange
		longLabel := ""
		for len(longLabel) <= 63 {
			longLabel += "a"
		}
		email := "test@" + longLabel + ".com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain label exceeds maximum length of 63 characters", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("top level domain too short", func(t *testing.T) {
		// Arrange
		email := "test@example.c"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email top-level domain must be at least 2 characters", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("top level domain with numbers", func(t *testing.T) {
		// Arrange
		email := "test@example.c0m"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email top-level domain must contain only letters", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain label starts with hyphen", func(t *testing.T) {
		// Arrange
		email := "test@-sub.example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain label cannot start or end with hyphen", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})

	t.Run("domain label ends with hyphen", func(t *testing.T) {
		// Arrange
		email := "test@sub-.example.com"

		// Act
		result, err := CreateEmailModel(email)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "email domain label cannot start or end with hyphen", err.Error())
		assert.Equal(t, EmailModel{}, result)
	})
}

func TestEmailModel_String(t *testing.T) {
	t.Run("returns email value", func(t *testing.T) {
		// Arrange
		expectedEmail := "test@example.com"
		emailModel, err := CreateEmailModel(expectedEmail)

		// Act
		result := emailModel.String()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedEmail, result)
	})
}
