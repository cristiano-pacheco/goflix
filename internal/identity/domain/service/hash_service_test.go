package service_test

import (
	"testing"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/service"
	"github.com/stretchr/testify/suite"
)

type HashServiceTestSuite struct {
	suite.Suite
	hashService service.HashService
}

func (suite *HashServiceTestSuite) SetupTest() {
	suite.hashService = service.NewHashService()
}

func TestHashServiceSuite(t *testing.T) {
	suite.Run(t, new(HashServiceTestSuite))
}

func (suite *HashServiceTestSuite) TestGenerateFromPassword() {
	// Arrange
	password := []byte("strongPassword123!")

	// Act
	hash, err := suite.hashService.GenerateFromPassword(password)

	// Assert
	suite.Require().NoError(err)
	suite.NotNil(hash)
	suite.NotEqual(password, hash)
	suite.Greater(len(hash), len(password))
}

func (suite *HashServiceTestSuite) TestCompareHashAndPassword_ValidPassword() {
	// Arrange
	password := []byte("strongPassword123!")
	hash, _ := suite.hashService.GenerateFromPassword(password)

	// Act
	err := suite.hashService.CompareHashAndPassword(hash, password)

	// Assert
	suite.NoError(err)
}

func (suite *HashServiceTestSuite) TestCompareHashAndPassword_InvalidPassword() {
	// Arrange
	correctPassword := []byte("strongPassword123!")
	wrongPassword := []byte("wrongPassword123!")
	hash, _ := suite.hashService.GenerateFromPassword(correctPassword)

	// Act
	err := suite.hashService.CompareHashAndPassword(hash, wrongPassword)

	// Assert
	suite.Error(err)
}

func (suite *HashServiceTestSuite) TestGenerateRandomBytes() {
	// Act
	randomBytes1, err1 := suite.hashService.GenerateRandomBytes()
	randomBytes2, err2 := suite.hashService.GenerateRandomBytes()

	// Assert
	suite.NoError(err1)
	suite.NoError(err2)
	suite.NotNil(randomBytes1)
	suite.NotNil(randomBytes2)
	suite.Len(randomBytes1, 128)               // Check default size
	suite.Len(randomBytes2, 128)               // Check default size
	suite.NotEqual(randomBytes1, randomBytes2) // Should be different random values
}

func (suite *HashServiceTestSuite) TestGenerateFromPassword_EmptyPassword() {
	// Arrange
	emptyPassword := []byte("")

	// Act
	hash, err := suite.hashService.GenerateFromPassword(emptyPassword)

	// Assert
	// Note: bcrypt actually allows empty passwords, so we expect no error
	suite.Require().NoError(err)
	suite.NotNil(hash)

	// Verify we can compare the empty password with its hash
	compareErr := suite.hashService.CompareHashAndPassword(hash, emptyPassword)
	suite.Require().NoError(compareErr)
}

func (suite *HashServiceTestSuite) TestCompareHashAndPassword_EmptyHash() {
	// Arrange
	password := []byte("strongPassword123!")
	emptyHash := []byte("")

	// Act
	err := suite.hashService.CompareHashAndPassword(emptyHash, password)

	// Assert
	suite.Error(err)
}
