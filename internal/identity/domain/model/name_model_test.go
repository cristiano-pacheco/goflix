package model_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
)

func TestCreateNameModel(t *testing.T) {
	t.Run("valid names", func(t *testing.T) {
		// Arrange
		validNames := []string{
			"John",
			"John Doe",
			"José",
			"François",
			"María García",
			"O'Connor",
			"Jean-Pierre",
			"Dr. Smith",
			"Anne-Marie",
			"João",
			"Müller",
			"李小明",
			"محمد",
			"Владимир",
		}

		for _, name := range validNames {
			// Act
			result, err := model.CreateNameModel(name)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, name, result.String())
		}
	})

	t.Run("valid name with whitespace trimming", func(t *testing.T) {
		// Arrange
		nameWithSpaces := "  John Doe  "
		expectedName := "John Doe"

		// Act
		result, err := model.CreateNameModel(nameWithSpaces)

		// Assert
		require.NoError(t, err)
		require.Equal(t, expectedName, result.String())
	})

	t.Run("empty name", func(t *testing.T) {
		// Arrange
		name := ""

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "name is required", err.Error())
		assert.Equal(t, model.NameModel{}, result)
	})

	t.Run("whitespace only name", func(t *testing.T) {
		// Arrange
		name := "   "

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "name is required", err.Error())
		assert.Equal(t, model.NameModel{}, result)
	})

	t.Run("name too short", func(t *testing.T) {
		// Arrange
		name := "J"

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.Error(t, err)
		require.Equal(t, "name must be at least 2 characters long", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("unicode name too short", func(t *testing.T) {
		// Arrange
		name := "李"

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.Error(t, err)
		require.Equal(t, "name must be at least 2 characters long", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name exceeds maximum length", func(t *testing.T) {
		// Arrange
		longName := strings.Repeat("a", 256)

		// Act
		result, err := model.CreateNameModel(longName)

		// Assert
		require.Error(t, err)
		require.Equal(t, "name cannot exceed 255 characters", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("unicode name exceeds maximum length", func(t *testing.T) {
		// Arrange
		longName := strings.Repeat("ç", 256)

		// Act
		result, err := model.CreateNameModel(longName)

		// Assert
		require.Error(t, err)
		require.Equal(t, "name cannot exceed 255 characters", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name starts with number", func(t *testing.T) {
		// Arrange
		name := "1John"

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.Error(t, err)
		require.Equal(t, "name must start with a letter", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name starts with hyphen", func(t *testing.T) {
		// Arrange
		name := "-John"

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.Error(t, err)
		require.Equal(t, "name must start with a letter", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name ends with hyphen", func(t *testing.T) {
		// Arrange
		name := "John-"

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.Error(t, err)
		require.Equal(t, "name must end with a letter or digit", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name with consecutive spaces", func(t *testing.T) {
		// Arrange
		name := "John  Doe"

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.Error(t, err)
		require.Equal(t, "name cannot contain consecutive spaces", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name with invalid characters", func(t *testing.T) {
		invalidNames := []string{
			"John@Doe",
			"John#Doe",
			"John$Doe",
			"John%Doe",
			"John&Doe",
			"John*Doe",
			"John+Doe",
			"John=Doe",
			"John?Doe",
			"John^Doe",
			"John_Doe",
			"John|Doe",
			"John~Doe",
		}

		for _, name := range invalidNames {
			// Arrange & Act
			result, err := model.CreateNameModel(name)

			// Assert
			require.Error(t, err)
			expectedError := "name contains invalid characters (only letters, digits, spaces, " +
				"hyphens, apostrophes, and periods are allowed)"
			require.Equal(t, expectedError, err.Error())
			require.Equal(t, model.NameModel{}, result)
		}
	})

	t.Run("name with excessive punctuation", func(t *testing.T) {
		// Arrange
		name := "John----Doe"

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.Error(t, err)
		require.Equal(t, "name cannot contain more than 3 consecutive punctuation marks", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name with leading spaces after trimming", func(t *testing.T) {
		// Arrange
		name := " John"

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.NoError(t, err)
		require.Equal(t, "John", result.String())
	})

	t.Run("name with trailing spaces after trimming", func(t *testing.T) {
		// Arrange
		name := "John "

		// Act
		result, err := model.CreateNameModel(name)

		// Assert
		require.NoError(t, err)
		require.Equal(t, "John", result.String())
	})
}

func TestNameModel_String(t *testing.T) {
	t.Run("returns name value", func(t *testing.T) {
		// Arrange
		expectedName := "John Doe"
		nameModel, err := model.CreateNameModel(expectedName)

		// Act
		result := nameModel.String()

		// Assert
		require.NoError(t, err)
		require.Equal(t, expectedName, result)
	})

	t.Run("returns unicode name value", func(t *testing.T) {
		// Arrange
		expectedName := "José María"
		nameModel, err := model.CreateNameModel(expectedName)

		// Act
		result := nameModel.String()

		// Assert
		require.NoError(t, err)
		require.Equal(t, expectedName, result)
	})
}
