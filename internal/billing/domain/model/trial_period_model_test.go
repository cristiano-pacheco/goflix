package model_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
)

func TestCreateTrialPeriodModel(t *testing.T) {
	t.Run("valid trial period returns model", func(t *testing.T) {
		// Arrange
		value := uint(7)

		// Act
		result, err := model.CreateTrialPeriodModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.Days())
	})

	t.Run("minimum trial period is valid", func(t *testing.T) {
		// Arrange
		value := uint(1)

		// Act
		result, err := model.CreateTrialPeriodModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.Days())
	})

	t.Run("maximum trial period is valid", func(t *testing.T) {
		// Arrange
		value := uint(30)

		// Act
		result, err := model.CreateTrialPeriodModel(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.Days())
	})

	t.Run("trial period below minimum returns error", func(t *testing.T) {
		// Arrange
		value := uint(0)

		// Act
		result, err := model.CreateTrialPeriodModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "trial period must be at least 1 day", err.Error())
		require.Equal(t, model.TrialPeriodModel{}, result)
	})

	t.Run("trial period above maximum returns error", func(t *testing.T) {
		// Arrange
		value := uint(31)

		// Act
		result, err := model.CreateTrialPeriodModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "trial period cannot exceed 365 days", err.Error())
		require.Equal(t, model.TrialPeriodModel{}, result)
	})

	t.Run("trial period way above maximum returns error", func(t *testing.T) {
		// Arrange
		value := uint(366)

		// Act
		result, err := model.CreateTrialPeriodModel(value)

		// Assert
		require.Error(t, err)
		require.Equal(t, "trial period cannot exceed 365 days", err.Error())
		require.Equal(t, model.TrialPeriodModel{}, result)
	})

	t.Run("common trial periods are valid", func(t *testing.T) {
		// Arrange
		commonPeriods := []uint{7, 14, 15, 30}

		for _, period := range commonPeriods {
			// Act
			result, err := model.CreateTrialPeriodModel(period)

			// Assert
			require.NoError(t, err)
			require.Equal(t, period, result.Days())
		}
	})
}

func TestTrialPeriodModel_Days(t *testing.T) {
	t.Run("returns the stored value", func(t *testing.T) {
		// Arrange
		value := uint(15)
		trialPeriod, err := model.CreateTrialPeriodModel(value)
		require.NoError(t, err)

		// Act
		result := trialPeriod.Days()

		// Assert
		require.Equal(t, value, result)
	})

	t.Run("returns minimum value correctly", func(t *testing.T) {
		// Arrange
		value := uint(1)
		trialPeriod, err := model.CreateTrialPeriodModel(value)
		require.NoError(t, err)

		// Act
		result := trialPeriod.Days()

		// Assert
		require.Equal(t, value, result)
	})

	t.Run("returns maximum value correctly", func(t *testing.T) {
		// Arrange
		value := uint(30)
		trialPeriod, err := model.CreateTrialPeriodModel(value)
		require.NoError(t, err)

		// Act
		result := trialPeriod.Days()

		// Assert
		require.Equal(t, value, result)
	})
}
