package model_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
)

func TestCreateNameModel(t *testing.T) {
	t.Run("valid name returns model", func(t *testing.T) {
		// Arrange
		value := "Basic Plan"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("name with leading and trailing spaces gets trimmed", func(t *testing.T) {
		// Arrange
		value := "  Premium Plan  "
		expected := "Premium Plan"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, expected, result.String())
	})

	t.Run("name starting with digit is valid", func(t *testing.T) {
		// Arrange
		value := "5G Plan"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("name with valid punctuation is valid", func(t *testing.T) {
		// Arrange
		value := "Pro-Plan_v2.0"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("empty name returns error", func(t *testing.T) {
		// Arrange
		value := ""

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name is required", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name with only spaces returns error", func(t *testing.T) {
		// Arrange
		value := "   "

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name is required", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name shorter than minimum length returns error", func(t *testing.T) {
		// Arrange
		value := "A"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name must be at least 2 characters long", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name at minimum length is valid", func(t *testing.T) {
		// Arrange
		value := "AB"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("name exceeding max length returns error", func(t *testing.T) {
		// Arrange
		value := strings.Repeat("a", 101)

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name cannot exceed 100 characters", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name at max length is valid", func(t *testing.T) {
		// Arrange
		value := strings.Repeat("a", 100)

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("name starting with punctuation returns error", func(t *testing.T) {
		// Arrange
		value := "-Invalid Plan"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name must start with a letter or digit", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name ending with punctuation returns error", func(t *testing.T) {
		// Arrange
		value := "Invalid Plan-"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name must end with a letter or digit", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name with consecutive spaces returns error", func(t *testing.T) {
		// Arrange
		value := "Invalid  Plan"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name cannot contain consecutive spaces", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name with invalid characters returns error", func(t *testing.T) {
		// Arrange
		value := "Invalid@Plan"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Contains(t, err.Error(), "plan name contains invalid characters")
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name with excessive consecutive punctuation returns error", func(t *testing.T) {
		// Arrange
		value := "Plan---Name"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name cannot contain more than 2 consecutive punctuation marks", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name with two consecutive punctuation marks is valid", func(t *testing.T) {
		// Arrange
		value := "Plan--Name"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("name with mixed valid punctuation is valid", func(t *testing.T) {
		// Arrange
		value := "Plan-Name_v1.0"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("name with unicode letters is valid", func(t *testing.T) {
		// Arrange
		value := "Plano Básico"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("name starting with underscore returns error", func(t *testing.T) {
		// Arrange
		value := "_Invalid Plan"

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name must start with a letter or digit", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})

	t.Run("name ending with period returns error", func(t *testing.T) {
		// Arrange
		value := "Invalid Plan."

		// Act
		result, err := model.CreateNameModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "plan name must end with a letter or digit", err.Error())
		require.Equal(t, model.NameModel{}, result)
	})
}

func TestNameModel_String(t *testing.T) {
	t.Run("returns the stored value", func(t *testing.T) {
		// Arrange
		value := "Test Plan"
		name, err := model.CreateNameModel(value)
		require.NoError(t, err)

		// Act
		result := name.String()

		// Assert
		require.Equal(t, value, result)
	})

	t.Run("returns trimmed value", func(t *testing.T) {
		// Arrange
		value := "  Trimmed Plan  "
		expected := "Trimmed Plan"
		name, err := model.CreateNameModel(value)
		require.NoError(t, err)

		// Act
		result := name.String()

		// Assert
		require.Equal(t, expected, result)
	})
}
