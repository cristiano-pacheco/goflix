package service_test

import (
	"testing"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/service"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), hash)
	assert.NotEqual(suite.T(), password, hash)
	assert.Greater(suite.T(), len(hash), len(password))
}

func (suite *HashServiceTestSuite) TestCompareHashAndPassword_ValidPassword() {
	// Arrange
	password := []byte("strongPassword123!")
	hash, _ := suite.hashService.GenerateFromPassword(password)

	// Act
	err := suite.hashService.CompareHashAndPassword(hash, password)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *HashServiceTestSuite) TestCompareHashAndPassword_InvalidPassword() {
	// Arrange
	correctPassword := []byte("strongPassword123!")
	wrongPassword := []byte("wrongPassword123!")
	hash, _ := suite.hashService.GenerateFromPassword(correctPassword)

	// Act
	err := suite.hashService.CompareHashAndPassword(hash, wrongPassword)

	// Assert
	assert.Error(suite.T(), err)
}

func (suite *HashServiceTestSuite) TestGenerateRandomBytes() {
	// Act
	randomBytes1, err1 := suite.hashService.GenerateRandomBytes()
	randomBytes2, err2 := suite.hashService.GenerateRandomBytes()

	// Assert
	assert.NoError(suite.T(), err1)
	assert.NoError(suite.T(), err2)
	assert.NotNil(suite.T(), randomBytes1)
	assert.NotNil(suite.T(), randomBytes2)
	assert.Equal(suite.T(), 128, len(randomBytes1))        // Check default size
	assert.Equal(suite.T(), 128, len(randomBytes2))        // Check default size
	assert.NotEqual(suite.T(), randomBytes1, randomBytes2) // Should be different random values
}

func (suite *HashServiceTestSuite) TestGenerateFromPassword_EmptyPassword() {
	// Arrange
	emptyPassword := []byte("")

	// Act
	hash, err := suite.hashService.GenerateFromPassword(emptyPassword)

	// Assert
	// Note: bcrypt actually allows empty passwords, so we expect no error
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), hash)

	// Verify we can compare the empty password with its hash
	compareErr := suite.hashService.CompareHashAndPassword(hash, emptyPassword)
	assert.NoError(suite.T(), compareErr)
}

func (suite *HashServiceTestSuite) TestCompareHashAndPassword_EmptyHash() {
	// Arrange
	password := []byte("strongPassword123!")
	emptyHash := []byte("")

	// Act
	err := suite.hashService.CompareHashAndPassword(emptyHash, password)

	// Assert
	assert.Error(suite.T(), err)
}
