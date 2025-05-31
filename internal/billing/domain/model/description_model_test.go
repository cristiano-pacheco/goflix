package model_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
)

func TestCreateDescriptionModel(t *testing.T) {
	t.Run("valid description returns model", func(t *testing.T) {
		// Arrange
		value := "This is a valid description"

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("empty description returns model", func(t *testing.T) {
		// Arrange
		value := ""

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, "", result.String())
	})

	t.Run("description with leading and trailing spaces gets trimmed", func(t *testing.T) {
		// Arrange
		value := "  valid description  "
		expected := "valid description"

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, expected, result.String())
	})

	t.Run("description with newlines and tabs is valid", func(t *testing.T) {
		// Arrange
		value := "Line 1\nLine 2\tTabbed"

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("description exceeding max length returns error", func(t *testing.T) {
		// Arrange
		value := strings.Repeat("a", 1001)

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "description cannot exceed 255 characters", err.Error())
		require.Equal(t, model.DescriptionModel{}, result)
	})

	t.Run("description at max length is valid", func(t *testing.T) {
		// Arrange
		value := strings.Repeat("a", 1000)

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("description with control characters returns error", func(t *testing.T) {
		// Arrange
		value := "Invalid\x00control"

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "description contains invalid characters (only printable characters are allowed)", err.Error())
		require.Equal(t, model.DescriptionModel{}, result)
	})

	t.Run("description with excessive consecutive spaces returns error", func(t *testing.T) {
		// Arrange
		value := "Too   many   spaces"

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "description cannot contain more than 2 consecutive spaces", err.Error())
		require.Equal(t, model.DescriptionModel{}, result)
	})

	t.Run("description with two consecutive spaces is valid", func(t *testing.T) {
		// Arrange
		value := "Two  spaces"

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("description with unicode characters is valid", func(t *testing.T) {
		// Arrange
		value := "Description with émojis 🎬 and açcénts"

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("description with non-printable characters returns error", func(t *testing.T) {
		// Arrange
		value := "Invalid\x7Fcharacter"

		// Act
		result, err := model.CreateDescriptionModel(value)

		// Assert
		require.Error(t, err)
		require.Contains(t, err.Error(), "description contains invalid characters")
		require.Equal(t, model.DescriptionModel{}, result)
	})
}

func TestDescriptionModel_String(t *testing.T) {
	t.Run("returns the stored value", func(t *testing.T) {
		// Arrange
		value := "Test description"
		description, err := model.CreateDescriptionModel(value)
		require.NoError(t, err)

		// Act
		result := description.String()

		// Assert
		require.Equal(t, value, result)
	})

	t.Run("returns empty string for empty description", func(t *testing.T) {
		// Arrange
		description, err := model.CreateDescriptionModel("")
		require.NoError(t, err)

		// Act
		result := description.String()

		// Assert
		require.Equal(t, "", result)
	})
}
