package model_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
)

func TestCreatePlanModel(t *testing.T) {
	t.Run("valid plan with all fields returns model", func(t *testing.T) {
		// Arrange
		name := "Premium Plan"
		description := "Premium subscription plan"
		currency := "USD"
		interval := "Month"
		amountCents := uint(2999)
		trialPeriod := uint(7)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, &trialPeriod)

		// Assert
		require.NoError(t, err)
		nameModel := result.Name()
		require.Equal(t, name, (&nameModel).String())
		require.NotNil(t, result.Description())
		require.Equal(t, description, result.Description().String())
		amountModel := result.Amount()
		require.Equal(t, amountCents, (&amountModel).Cents())
		currencyModel := result.Currency()
		require.Equal(t, currency, currencyModel.Code())
		intervalModel := result.Interval()
		require.Equal(t, interval, intervalModel.String())
		require.NotNil(t, result.TrialPeriod())
		require.Equal(t, trialPeriod, result.TrialPeriod().Days())
		require.True(t, result.CreatedAt().After(time.Time{}))
		require.True(t, result.UpdatedAt().After(time.Time{}))
		require.Equal(t, uint64(0), result.ID())
	})

	t.Run("valid plan without description returns model", func(t *testing.T) {
		// Arrange
		name := "Basic Plan"
		description := ""
		currency := "USD"
		interval := "Month"
		amountCents := uint(999)
		trialPeriod := uint(14)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, &trialPeriod)

		// Assert
		require.NoError(t, err)
		nameModel := result.Name()
		require.Equal(t, name, (&nameModel).String())
		require.Nil(t, result.Description())
		amountModel := result.Amount()
		require.Equal(t, amountCents, (&amountModel).Cents())
		currencyModel := result.Currency()
		require.Equal(t, currency, currencyModel.Code())
		intervalModel := result.Interval()
		require.Equal(t, interval, intervalModel.String())
		require.NotNil(t, result.TrialPeriod())
		require.Equal(t, trialPeriod, result.TrialPeriod().Days())
	})

	t.Run("valid plan without trial period returns model", func(t *testing.T) {
		// Arrange
		name := "Enterprise Plan"
		description := "Enterprise subscription plan"
		currency := "EUR"
		interval := "Year"
		amountCents := uint(99999)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, nil)

		// Assert
		require.NoError(t, err)
		nameModel := result.Name()
		require.Equal(t, name, (&nameModel).String())
		require.NotNil(t, result.Description())
		require.Equal(t, description, result.Description().String())
		amountModel := result.Amount()
		require.Equal(t, amountCents, (&amountModel).Cents())
		currencyModel := result.Currency()
		require.Equal(t, currency, currencyModel.Code())
		intervalModel := result.Interval()
		require.Equal(t, interval, intervalModel.String())
		require.Nil(t, result.TrialPeriod())
	})

	t.Run("plan with leading and trailing spaces gets trimmed", func(t *testing.T) {
		// Arrange
		name := "  Trimmed Plan  "
		description := "  Trimmed description  "
		currency := "  USD  "
		interval := "Month"
		amountCents := uint(1999)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, nil)

		// Assert
		require.NoError(t, err)
		nameModel := result.Name()
		require.Equal(t, "Trimmed Plan", (&nameModel).String())
		require.Equal(t, "Trimmed description", result.Description().String())
		currencyModel := result.Currency()
		require.Equal(t, "USD", currencyModel.Code())
	})

	t.Run("plan with different valid intervals", func(t *testing.T) {
		// Arrange
		validIntervals := []string{"Day", "Week", "Month", "Year"}
		name := "Test Plan"
		currency := "USD"
		amountCents := uint(999)

		for _, interval := range validIntervals {
			// Act
			result, err := model.CreatePlanModel(name, "", currency, interval, amountCents, nil)

			// Assert
			require.NoError(t, err)
			intervalModel := result.Interval()
			require.Equal(t, interval, intervalModel.String())
		}
	})

	t.Run("invalid name returns error", func(t *testing.T) {
		// Arrange
		name := "A"
		description := "Valid description"
		currency := "USD"
		interval := "Month"
		amountCents := uint(999)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, nil)

		// Assert
		require.Error(t, err)
		require.Contains(t, err.Error(), "plan name must be at least 2 characters long")
		require.Equal(t, model.PlanModel{}, result)
	})

	t.Run("invalid description returns error", func(t *testing.T) {
		// Arrange
		name := "Valid Plan"
		description := "Invalid\x00description"
		currency := "USD"
		interval := "Month"
		amountCents := uint(999)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, nil)

		// Assert
		require.Error(t, err)
		require.Contains(
			t,
			err.Error(),
			"description contains invalid characters (only printable characters are allowed)",
		)
		require.Equal(t, model.PlanModel{}, result)
	})

	t.Run("invalid amount returns error", func(t *testing.T) {
		// Arrange
		name := "Valid Plan"
		description := "Valid description"
		currency := "USD"
		interval := "Month"
		amountCents := uint(1000000000)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, nil)

		// Assert
		require.Error(t, err)
		require.Contains(t, err.Error(), "amount exceeds maximum allowed value")
		require.Equal(t, model.PlanModel{}, result)
	})

	t.Run("invalid currency returns error", func(t *testing.T) {
		// Arrange
		name := "Valid Plan"
		description := "Valid description"
		currency := "INVALID"
		interval := "Month"
		amountCents := uint(999)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, nil)

		// Assert
		require.Error(t, err)
		require.Contains(t, err.Error(), "currency code must be exactly 3 characters")
		require.Equal(t, model.PlanModel{}, result)
	})

	t.Run("invalid interval returns error", func(t *testing.T) {
		// Arrange
		name := "Valid Plan"
		description := "Valid description"
		currency := "USD"
		interval := "Invalid"
		amountCents := uint(999)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, nil)

		// Assert
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid plan interval")
		require.Equal(t, model.PlanModel{}, result)
	})

	t.Run("invalid trial period returns error", func(t *testing.T) {
		// Arrange
		name := "Valid Plan"
		description := "Valid description"
		currency := "USD"
		interval := "Month"
		amountCents := uint(999)
		trialPeriod := uint(0)

		// Act
		result, err := model.CreatePlanModel(name, description, currency, interval, amountCents, &trialPeriod)

		// Assert
		require.Error(t, err)
		require.Contains(t, err.Error(), "trial period must be at least 1 day")
		require.Equal(t, model.PlanModel{}, result)
	})
}

func TestRestorePlanModel(t *testing.T) {
	t.Run("valid plan restoration with all fields returns model", func(t *testing.T) {
		// Arrange
		id := uint64(123)
		name := "Restored Plan"
		description := "Restored description"
		currency := "USD"
		interval := "Month"
		amountCents := uint(1999)
		trialPeriod := uint(7)
		createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		updatedAt := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)

		// Act
		result, err := model.RestorePlanModel(
			id,
			name,
			description,
			currency,
			interval,
			amountCents,
			&trialPeriod,
			createdAt,
			updatedAt,
		)

		// Assert
		require.NoError(t, err)
		require.Equal(t, id, result.ID())
		nameModel := result.Name()
		require.Equal(t, name, (&nameModel).String())
		require.NotNil(t, result.Description())
		require.Equal(t, description, result.Description().String())
		amountModel := result.Amount()
		require.Equal(t, amountCents, (&amountModel).Cents())
		currencyModel := result.Currency()
		require.Equal(t, currency, currencyModel.Code())
		intervalModel := result.Interval()
		require.Equal(t, interval, intervalModel.String())
		require.NotNil(t, result.TrialPeriod())
		require.Equal(t, trialPeriod, result.TrialPeriod().Days())
		require.Equal(t, createdAt, result.CreatedAt())
		require.Equal(t, updatedAt, result.UpdatedAt())
	})

	t.Run("valid plan restoration without optional fields returns model", func(t *testing.T) {
		// Arrange
		id := uint64(456)
		name := "Basic Restored Plan"
		description := ""
		currency := "EUR"
		interval := "Year"
		amountCents := uint(9999)
		createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		updatedAt := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)

		// Act
		result, err := model.RestorePlanModel(
			id,
			name,
			description,
			currency,
			interval,
			amountCents,
			nil,
			createdAt,
			updatedAt,
		)

		// Assert
		require.NoError(t, err)
		require.Equal(t, id, result.ID())
		nameModel := result.Name()
		require.Equal(t, name, (&nameModel).String())
		require.Nil(t, result.Description())
		amountModel := result.Amount()
		require.Equal(t, amountCents, (&amountModel).Cents())
		currencyModel := result.Currency()
		require.Equal(t, currency, currencyModel.Code())
		intervalModel := result.Interval()
		require.Equal(t, interval, intervalModel.String())
		require.Nil(t, result.TrialPeriod())
		require.Equal(t, createdAt, result.CreatedAt())
		require.Equal(t, updatedAt, result.UpdatedAt())
	})

	t.Run("invalid name in restoration returns error", func(t *testing.T) {
		// Arrange
		id := uint64(789)
		name := ""
		description := "Valid description"
		currency := "USD"
		interval := "Month"
		amountCents := uint(999)
		createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		updatedAt := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)

		// Act
		result, err := model.RestorePlanModel(
			id,
			name,
			description,
			currency,
			interval,
			amountCents,
			nil,
			createdAt,
			updatedAt,
		)

		// Assert
		require.Error(t, err)
		require.Contains(t, err.Error(), "plan name is required")
		require.Equal(t, model.PlanModel{}, result)
	})
}

func TestPlanModel_Getters(t *testing.T) {
	t.Run("all getters return correct values", func(t *testing.T) {
		// Arrange
		id := uint64(999)
		name := "Test Plan"
		description := "Test description"
		currency := "USD"
		interval := "Month"
		amountCents := uint(2999)
		trialPeriod := uint(14)
		createdAt := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		updatedAt := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)

		plan, err := model.RestorePlanModel(
			id,
			name,
			description,
			currency,
			interval,
			amountCents,
			&trialPeriod,
			createdAt,
			updatedAt,
		)
		require.NoError(t, err)

		// Act & Assert
		require.Equal(t, id, plan.ID())
		nameModel := plan.Name()
		require.Equal(t, name, (&nameModel).String())
		require.NotNil(t, plan.Description())
		require.Equal(t, description, plan.Description().String())
		amountModel := plan.Amount()
		require.Equal(t, amountCents, (&amountModel).Cents())
		currencyModel := plan.Currency()
		require.Equal(t, currency, currencyModel.Code())
		intervalModel := plan.Interval()
		require.Equal(t, interval, intervalModel.String())
		require.NotNil(t, plan.TrialPeriod())
		require.Equal(t, trialPeriod, plan.TrialPeriod().Days())
		require.Equal(t, createdAt, plan.CreatedAt())
		require.Equal(t, updatedAt, plan.UpdatedAt())
	})

	t.Run("getters return correct values for plan without optional fields", func(t *testing.T) {
		// Arrange
		name := "Minimal Plan"
		currency := "EUR"
		interval := "Year"
		amountCents := uint(9999)

		plan, err := model.CreatePlanModel(name, "", currency, interval, amountCents, nil)
		require.NoError(t, err)

		// Act & Assert
		require.Equal(t, uint64(0), plan.ID())
		nameModel := plan.Name()
		require.Equal(t, name, (&nameModel).String())
		require.Nil(t, plan.Description())
		amountModel := plan.Amount()
		require.Equal(t, amountCents, (&amountModel).Cents())
		currencyModel := plan.Currency()
		require.Equal(t, currency, currencyModel.Code())
		intervalModel := plan.Interval()
		require.Equal(t, interval, intervalModel.String())
		require.Nil(t, plan.TrialPeriod())
		require.True(t, plan.CreatedAt().After(time.Time{}))
		require.True(t, plan.UpdatedAt().After(time.Time{}))
	})
}
