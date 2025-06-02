package enum_test

import (
	"testing"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/enum"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
	"github.com/stretchr/testify/require"
)

func TestNewSubscriptionStatusEnum(t *testing.T) {
	t.Run("valid active status returns enum without error", func(t *testing.T) {
		// Arrange
		value := enum.EnumSubscriptionStatusActive

		// Act
		result, err := enum.NewSubscriptionStatusEnum(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("valid inactive status returns enum without error", func(t *testing.T) {
		// Arrange
		value := enum.EnumSubscriptionStatusInactive

		// Act
		result, err := enum.NewSubscriptionStatusEnum(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("valid cancelled status returns enum without error", func(t *testing.T) {
		// Arrange
		value := enum.EnumSubscriptionStatusCancelled

		// Act
		result, err := enum.NewSubscriptionStatusEnum(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("valid expired status returns enum without error", func(t *testing.T) {
		// Arrange
		value := enum.EnumSubscriptionStatusExpired

		// Act
		result, err := enum.NewSubscriptionStatusEnum(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("valid past due status returns enum without error", func(t *testing.T) {
		// Arrange
		value := enum.EnumSubscriptionStatusPastDue

		// Act
		result, err := enum.NewSubscriptionStatusEnum(value)

		// Assert
		require.NoError(t, err)
		require.Equal(t, value, result.String())
	})

	t.Run("invalid status returns error", func(t *testing.T) {
		// Arrange
		value := "InvalidStatus"

		// Act
		result, err := enum.NewSubscriptionStatusEnum(value)

		// Assert
		require.ErrorIs(t, err, errs.ErrInvalidSubscriptionStatus)
		require.Equal(t, "", result.String())
	})

	t.Run("empty string returns error", func(t *testing.T) {
		// Arrange
		value := ""

		// Act
		result, err := enum.NewSubscriptionStatusEnum(value)

		// Assert
		require.ErrorIs(t, err, errs.ErrInvalidSubscriptionStatus)
		require.Equal(t, "", result.String())
	})

	t.Run("case sensitive validation returns error for lowercase", func(t *testing.T) {
		// Arrange
		value := "active"

		// Act
		result, err := enum.NewSubscriptionStatusEnum(value)

		// Assert
		require.ErrorIs(t, err, errs.ErrInvalidSubscriptionStatus)
		require.Equal(t, "", result.String())
	})
}

func TestSubscriptionStatusEnum_String(t *testing.T) {
	t.Run("returns correct string value for active", func(t *testing.T) {
		// Arrange
		subscriptionStatus, err := enum.NewSubscriptionStatusEnum(enum.EnumSubscriptionStatusActive)
		require.NoError(t, err)

		// Act
		result := subscriptionStatus.String()

		// Assert
		require.Equal(t, enum.EnumSubscriptionStatusActive, result)
	})

	t.Run("returns correct string value for cancelled", func(t *testing.T) {
		// Arrange
		subscriptionStatus, err := enum.NewSubscriptionStatusEnum(enum.EnumSubscriptionStatusCancelled)
		require.NoError(t, err)

		// Act
		result := subscriptionStatus.String()

		// Assert
		require.Equal(t, enum.EnumSubscriptionStatusCancelled, result)
	})
} 