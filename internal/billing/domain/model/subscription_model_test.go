package model_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/enum"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
)

func TestCreateSubscriptionModel(t *testing.T) {
	t.Run("valid input returns subscription model", func(t *testing.T) {
		// Arrange
		userID := uint64(1)
		planID := uint64(2)
		status := enum.EnumSubscriptionStatusActive
		startDate := time.Now().UTC()
		endDate := &time.Time{}
		*endDate = startDate.Add(30 * 24 * time.Hour)
		autoRenew := true

		// Act
		subscription, err := model.CreateSubscriptionModel(userID, planID, status, startDate, endDate, autoRenew)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, userID, subscription.UserID())
		assert.Equal(t, planID, subscription.PlanID())
		statusEnum := subscription.Status()
		assert.Equal(t, status, (&statusEnum).String())
		assert.Equal(t, startDate, subscription.StartDate())
		assert.Equal(t, endDate, subscription.EndDate())
		assert.Equal(t, autoRenew, subscription.AutoRenew())
		assert.NotZero(t, subscription.CreatedAt())
		assert.NotZero(t, subscription.UpdatedAt())
		assert.Zero(t, subscription.ID())
	})

	t.Run("valid input with nil end date returns subscription model", func(t *testing.T) {
		// Arrange
		userID := uint64(1)
		planID := uint64(2)
		status := enum.EnumSubscriptionStatusActive
		startDate := time.Now().UTC()
		autoRenew := false

		// Act
		subscription, err := model.CreateSubscriptionModel(userID, planID, status, startDate, nil, autoRenew)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, userID, subscription.UserID())
		assert.Equal(t, planID, subscription.PlanID())
		statusEnum := subscription.Status()
		assert.Equal(t, status, (&statusEnum).String())
		assert.Equal(t, startDate, subscription.StartDate())
		assert.Nil(t, subscription.EndDate())
		assert.Equal(t, autoRenew, subscription.AutoRenew())
	})

	t.Run("zero user ID returns error", func(t *testing.T) {
		// Arrange
		userID := uint64(0)
		planID := uint64(2)
		status := enum.EnumSubscriptionStatusActive
		startDate := time.Now().UTC()
		autoRenew := true

		// Act
		subscription, err := model.CreateSubscriptionModel(userID, planID, status, startDate, nil, autoRenew)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "user ID is required", err.Error())
		assert.Zero(t, subscription)
	})

	t.Run("zero plan ID returns error", func(t *testing.T) {
		// Arrange
		userID := uint64(1)
		planID := uint64(0)
		status := enum.EnumSubscriptionStatusActive
		startDate := time.Now().UTC()
		autoRenew := true

		// Act
		subscription, err := model.CreateSubscriptionModel(userID, planID, status, startDate, nil, autoRenew)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "plan ID is required", err.Error())
		assert.Zero(t, subscription)
	})

	t.Run("zero start date returns error", func(t *testing.T) {
		// Arrange
		userID := uint64(1)
		planID := uint64(2)
		status := enum.EnumSubscriptionStatusActive
		startDate := time.Time{}
		autoRenew := true

		// Act
		subscription, err := model.CreateSubscriptionModel(userID, planID, status, startDate, nil, autoRenew)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "start date is required", err.Error())
		assert.Zero(t, subscription)
	})

	t.Run("end date before start date returns error", func(t *testing.T) {
		// Arrange
		userID := uint64(1)
		planID := uint64(2)
		status := enum.EnumSubscriptionStatusActive
		startDate := time.Now().UTC()
		endDate := &time.Time{}
		*endDate = startDate.Add(-24 * time.Hour)
		autoRenew := true

		// Act
		subscription, err := model.CreateSubscriptionModel(userID, planID, status, startDate, endDate, autoRenew)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "end date cannot be before start date", err.Error())
		assert.Zero(t, subscription)
	})

	t.Run("invalid status returns error", func(t *testing.T) {
		// Arrange
		userID := uint64(1)
		planID := uint64(2)
		status := "invalid_status"
		startDate := time.Now().UTC()
		autoRenew := true

		// Act
		subscription, err := model.CreateSubscriptionModel(userID, planID, status, startDate, nil, autoRenew)

		// Assert
		require.Error(t, err)
		assert.Zero(t, subscription)
	})
}

func TestRestoreSubscriptionModel(t *testing.T) {
	t.Run("valid input returns subscription model", func(t *testing.T) {
		// Arrange
		id := uint64(10)
		userID := uint64(1)
		planID := uint64(2)
		status := enum.EnumSubscriptionStatusActive
		startDate := time.Now().UTC()
		endDate := &time.Time{}
		*endDate = startDate.Add(30 * 24 * time.Hour)
		autoRenew := true
		createdAt := time.Now().UTC().Add(-24 * time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		subscription, err := model.RestoreSubscriptionModel(
			id, userID, planID, status, startDate, endDate, autoRenew, createdAt, updatedAt,
		)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, id, subscription.ID())
		assert.Equal(t, userID, subscription.UserID())
		assert.Equal(t, planID, subscription.PlanID())
		statusEnum := subscription.Status()
		assert.Equal(t, status, (&statusEnum).String())
		assert.Equal(t, startDate, subscription.StartDate())
		assert.Equal(t, endDate, subscription.EndDate())
		assert.Equal(t, autoRenew, subscription.AutoRenew())
		assert.Equal(t, createdAt, subscription.CreatedAt())
		assert.Equal(t, updatedAt, subscription.UpdatedAt())
	})

	t.Run("zero user ID returns error", func(t *testing.T) {
		// Arrange
		id := uint64(10)
		userID := uint64(0)
		planID := uint64(2)
		status := enum.EnumSubscriptionStatusActive
		startDate := time.Now().UTC()
		autoRenew := true
		createdAt := time.Now().UTC().Add(-24 * time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		subscription, err := model.RestoreSubscriptionModel(
			id, userID, planID, status, startDate, nil, autoRenew, createdAt, updatedAt,
		)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "user ID is required", err.Error())
		assert.Zero(t, subscription)
	})

	t.Run("invalid status returns error", func(t *testing.T) {
		// Arrange
		id := uint64(10)
		userID := uint64(1)
		planID := uint64(2)
		status := "invalid_status"
		startDate := time.Now().UTC()
		autoRenew := true
		createdAt := time.Now().UTC().Add(-24 * time.Hour)
		updatedAt := time.Now().UTC()

		// Act
		subscription, err := model.RestoreSubscriptionModel(
			id, userID, planID, status, startDate, nil, autoRenew, createdAt, updatedAt,
		)

		// Assert
		require.Error(t, err)
		assert.Zero(t, subscription)
	})
}

func TestSubscriptionModel_UpdateStatus(t *testing.T) {
	t.Run("valid status updates successfully", func(t *testing.T) {
		// Arrange
		subscription, err := model.CreateSubscriptionModel(
			1, 2, enum.EnumSubscriptionStatusActive, time.Now().UTC(), nil, true,
		)
		require.NoError(t, err)
		originalUpdatedAt := subscription.UpdatedAt()
		time.Sleep(time.Millisecond)

		// Act
		err = subscription.UpdateStatus(enum.EnumSubscriptionStatusCancelled)

		// Assert
		require.NoError(t, err)
		statusEnum := subscription.Status()
		assert.Equal(t, enum.EnumSubscriptionStatusCancelled, (&statusEnum).String())
		assert.True(t, subscription.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("invalid status returns error", func(t *testing.T) {
		// Arrange
		subscription, err := model.CreateSubscriptionModel(
			1, 2, enum.EnumSubscriptionStatusActive, time.Now().UTC(), nil, true,
		)
		require.NoError(t, err)
		originalStatus := subscription.Status()
		originalUpdatedAt := subscription.UpdatedAt()

		// Act
		err = subscription.UpdateStatus("invalid_status")

		// Assert
		require.Error(t, err)
		assert.Equal(t, originalStatus, subscription.Status())
		assert.Equal(t, originalUpdatedAt, subscription.UpdatedAt())
	})
}

func TestSubscriptionModel_UpdateEndDate(t *testing.T) {
	t.Run("valid end date updates successfully", func(t *testing.T) {
		// Arrange
		startDate := time.Now().UTC()
		subscription, err := model.CreateSubscriptionModel(
			1, 2, enum.EnumSubscriptionStatusActive, startDate, nil, true,
		)
		require.NoError(t, err)
		originalUpdatedAt := subscription.UpdatedAt()
		time.Sleep(time.Millisecond)
		newEndDate := &time.Time{}
		*newEndDate = startDate.Add(60 * 24 * time.Hour)

		// Act
		err = subscription.UpdateEndDate(newEndDate)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, newEndDate, subscription.EndDate())
		assert.True(t, subscription.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("nil end date updates successfully", func(t *testing.T) {
		// Arrange
		startDate := time.Now().UTC()
		endDate := &time.Time{}
		*endDate = startDate.Add(30 * 24 * time.Hour)
		subscription, err := model.CreateSubscriptionModel(
			1, 2, enum.EnumSubscriptionStatusActive, startDate, endDate, true,
		)
		require.NoError(t, err)
		originalUpdatedAt := subscription.UpdatedAt()
		time.Sleep(time.Millisecond)

		// Act
		err = subscription.UpdateEndDate(nil)

		// Assert
		require.NoError(t, err)
		assert.Nil(t, subscription.EndDate())
		assert.True(t, subscription.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("end date before start date returns error", func(t *testing.T) {
		// Arrange
		startDate := time.Now().UTC()
		subscription, err := model.CreateSubscriptionModel(
			1, 2, enum.EnumSubscriptionStatusActive, startDate, nil, true,
		)
		require.NoError(t, err)
		originalEndDate := subscription.EndDate()
		originalUpdatedAt := subscription.UpdatedAt()
		invalidEndDate := &time.Time{}
		*invalidEndDate = startDate.Add(-24 * time.Hour)

		// Act
		err = subscription.UpdateEndDate(invalidEndDate)

		// Assert
		require.Error(t, err)
		assert.Equal(t, "end date cannot be before start date", err.Error())
		assert.Equal(t, originalEndDate, subscription.EndDate())
		assert.Equal(t, originalUpdatedAt, subscription.UpdatedAt())
	})
}

func TestSubscriptionModel_SetAutoRenew(t *testing.T) {
	t.Run("sets auto renew to true", func(t *testing.T) {
		// Arrange
		subscription, err := model.CreateSubscriptionModel(
			1, 2, enum.EnumSubscriptionStatusActive, time.Now().UTC(), nil, false,
		)
		require.NoError(t, err)
		originalUpdatedAt := subscription.UpdatedAt()
		time.Sleep(time.Millisecond)

		// Act
		subscription.SetAutoRenew(true)

		// Assert
		assert.True(t, subscription.AutoRenew())
		assert.True(t, subscription.UpdatedAt().After(originalUpdatedAt))
	})

	t.Run("sets auto renew to false", func(t *testing.T) {
		// Arrange
		subscription, err := model.CreateSubscriptionModel(
			1, 2, enum.EnumSubscriptionStatusActive, time.Now().UTC(), nil, true,
		)
		require.NoError(t, err)
		originalUpdatedAt := subscription.UpdatedAt()
		time.Sleep(time.Millisecond)

		// Act
		subscription.SetAutoRenew(false)

		// Assert
		assert.False(t, subscription.AutoRenew())
		assert.True(t, subscription.UpdatedAt().After(originalUpdatedAt))
	})
}
