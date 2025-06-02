package enum_test

import (
	"testing"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/enum"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
	"github.com/stretchr/testify/require"
)

func TestNewPlanIntervalEnum(t *testing.T) {
	t.Run("valid day interval returns enum without error", func(t *testing.T) {
		// Arrange
		value := enum.EnumPlanIntervalDay

		// Act
		result, err := enum.NewPlanIntervalEnum(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("valid week interval returns enum without error", func(t *testing.T) {
		// Arrange
		value := enum.EnumPlanIntervalWeek

		// Act
		result, err := enum.NewPlanIntervalEnum(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("valid month interval returns enum without error", func(t *testing.T) {
		// Arrange
		value := enum.EnumPlanIntervalMonth

		// Act
		result, err := enum.NewPlanIntervalEnum(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("valid year interval returns enum without error", func(t *testing.T) {
		// Arrange
		value := enum.EnumPlanIntervalYear

		// Act
		result, err := enum.NewPlanIntervalEnum(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("invalid interval returns error", func(t *testing.T) {
		// Arrange
		value := "InvalidInterval"

		// Act
		result, err := enum.NewPlanIntervalEnum(value)

		// Assert
		require.ErrorIs(t, err, errs.ErrInvalidPlanInterval)
		require.Equal(t, "", result.String())
	})

	t.Run("empty string returns error", func(t *testing.T) {
		// Arrange
		value := ""

		// Act
		result, err := enum.NewPlanIntervalEnum(value)

		// Assert
		require.ErrorIs(t, err, errs.ErrInvalidPlanInterval)
		require.Equal(t, "", result.String())
	})

	t.Run("case sensitive validation returns error for lowercase", func(t *testing.T) {
		// Arrange
		value := "day"

		// Act
		result, err := enum.NewPlanIntervalEnum(value)

		// Assert
		require.ErrorIs(t, err, errs.ErrInvalidPlanInterval)
		require.Equal(t, "", result.String())
	})
}

func TestPlanIntervalEnum_String(t *testing.T) {
	t.Run("returns correct string value for day", func(t *testing.T) {
		// Arrange
		planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalDay)
		require.NoError(t, err)

		// Act
		result := planInterval.String()

		// Assert
		require.Equal(t, enum.EnumPlanIntervalDay, result)
	})

	t.Run("returns correct string value for month", func(t *testing.T) {
		// Arrange
		planInterval, err := enum.NewPlanIntervalEnum(enum.EnumPlanIntervalMonth)
		require.NoError(t, err)

		// Act
		result := planInterval.String()

		// Assert
		require.Equal(t, enum.EnumPlanIntervalMonth, result)
	})
} 