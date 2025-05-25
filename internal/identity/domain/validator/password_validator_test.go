package validator_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	customerr "github.com/cristiano-pacheco/goflix/internal/identity/domain/errs"
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/validator"
)

type PasswordValidatorTestSuite struct {
	suite.Suite
	validator validator.PasswordValidator
}

func (suite *PasswordValidatorTestSuite) SetupTest() {
	suite.validator = validator.NewPasswordValidator()
}

func TestPasswordValidatorSuite(t *testing.T) {
	suite.Run(t, new(PasswordValidatorTestSuite))
}

func (suite *PasswordValidatorTestSuite) TestValidate_ValidPassword() {
	// Arrange
	validPassword := "Abc123!@"

	// Act
	err := suite.validator.Validate(validPassword)

	// Assert
	suite.NoError(err)
}

func (suite *PasswordValidatorTestSuite) TestValidate_TooShort() {
	// Arrange
	shortPassword := "Abc12!"

	// Act
	err := suite.validator.Validate(shortPassword)

	// Assert
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, customerr.ErrPasswordTooShort)
}

func (suite *PasswordValidatorTestSuite) TestValidate_NoUppercase() {
	// Arrange
	noUppercasePassword := "abc123!@"

	// Act
	err := suite.validator.Validate(noUppercasePassword)

	// Assert
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, customerr.ErrPasswordNoUppercase)
}

func (suite *PasswordValidatorTestSuite) TestValidate_NoLowercase() {
	// Arrange
	noLowercasePassword := "ABC123!@"

	// Act
	err := suite.validator.Validate(noLowercasePassword)

	// Assert
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, customerr.ErrPasswordNoLowercase)
}

func (suite *PasswordValidatorTestSuite) TestValidate_NoNumber() {
	// Arrange
	noNumberPassword := "Abcdef!@"

	// Act
	err := suite.validator.Validate(noNumberPassword)

	// Assert
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, customerr.ErrPasswordNoNumber)
}

func (suite *PasswordValidatorTestSuite) TestValidate_NoSpecialCharacter() {
	// Arrange
	noSpecialCharPassword := "Abcdef123"

	// Act
	err := suite.validator.Validate(noSpecialCharPassword)

	// Assert
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, customerr.ErrPasswordNoSpecialChar)
}

func (suite *PasswordValidatorTestSuite) TestValidate_UTF8Support() {
	// Arrange
	// Contains: uppercase, lowercase, number, special char, and UTF-8 characters
	utf8Password := "Пароль123!@" // Russian word for "password" with numbers and special chars

	// Act
	err := suite.validator.Validate(utf8Password)

	// Assert
	suite.NoError(err)
}

func (suite *PasswordValidatorTestSuite) TestValidate_UTF8Length() {
	// Arrange
	// This password has 6 characters but more bytes
	shortUtf8Password := "Парол!" // 6 characters (less than 8 required)

	// Act
	err := suite.validator.Validate(shortUtf8Password)

	// Assert
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, customerr.ErrPasswordTooShort)
}

func (suite *PasswordValidatorTestSuite) TestValidate_EmptyPassword() {
	// Arrange
	emptyPassword := ""

	// Act
	err := suite.validator.Validate(emptyPassword)

	// Assert
	suite.Require().Error(err)
	suite.Require().ErrorIs(err, customerr.ErrPasswordTooShort)
}

func (suite *PasswordValidatorTestSuite) TestValidate_ComplexUTF8Password() {
	// Arrange
	// Contains: Chinese characters (汉字), uppercase, lowercase, numbers, and special chars
	complexPassword := "汉字Abc123!@"

	// Act
	err := suite.validator.Validate(complexPassword)

	// Assert
	suite.NoError(err)
}

func (suite *PasswordValidatorTestSuite) TestValidate_EmojiPassword() {
	// Arrange
	// Contains: Emoji, uppercase, lowercase, numbers, and special chars
	emojiPassword := "Pass123!😀"

	// Act
	err := suite.validator.Validate(emojiPassword)

	// Assert
	suite.NoError(err)
}
