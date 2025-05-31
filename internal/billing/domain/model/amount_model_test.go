package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
)

func TestCreateAmountModel(t *testing.T) {
	t.Run("valid amount creates model successfully", func(t *testing.T) {
		// Arrange
		validAmount := uint(100)

		// Act
		result, err := model.CreateAmountModel(validAmount)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, validAmount, result.Cents())
	})

	t.Run("zero amount creates model successfully", func(t *testing.T) {
		// Arrange
		zeroAmount := uint(0)

		// Act
		result, err := model.CreateAmountModel(zeroAmount)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, zeroAmount, result.Cents())
	})

	t.Run("maximum allowed amount creates model successfully", func(t *testing.T) {
		// Arrange
		maxAmount := uint(999999999)

		// Act
		result, err := model.CreateAmountModel(maxAmount)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, maxAmount, result.Cents())
	})

	t.Run("amount exceeding maximum returns error", func(t *testing.T) {
		// Arrange
		exceedingAmount := uint(1000000000)

		// Act
		result, err := model.CreateAmountModel(exceedingAmount)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "amount exceeds maximum allowed value", err.Error())
		assert.Equal(t, model.AmountModel{}, result)
	})
}

func TestAmountModel_Cents(t *testing.T) {
	t.Run("returns correct cents value", func(t *testing.T) {
		// Arrange
		expectedCents := uint(12345)
		amountModel, err := model.CreateAmountModel(expectedCents)
		require.NoError(t, err)

		// Act
		result := amountModel.Cents()

		// Assert
		assert.Equal(t, expectedCents, result)
	})
}
