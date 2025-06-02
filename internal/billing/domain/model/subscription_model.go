package model

import (
	"time"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/enum"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
)

type SubscriptionModel struct {
	id        uint64
	userID    uint64
	planID    uint64
	status    enum.SubscriptionStatusEnum
	startDate time.Time
	endDate   *time.Time
	autoRenew bool
	createdAt time.Time
	updatedAt time.Time
}

func CreateSubscriptionModel(
	userID, planID uint64,
	startDate time.Time,
	endDate *time.Time,
) (SubscriptionModel, error) {
	if err := validateSubscription(userID, planID, startDate, endDate); err != nil {
		return SubscriptionModel{}, err
	}

	statusEnum, err := enum.NewSubscriptionStatusEnum(enum.EnumSubscriptionStatusActive)
	if err != nil {
		return SubscriptionModel{}, err
	}

	return SubscriptionModel{
		userID:    userID,
		planID:    planID,
		status:    statusEnum,
		startDate: startDate,
		endDate:   endDate,
		autoRenew: true, // auto-renew enabled by default
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}, nil
}

func RestoreSubscriptionModel(
	id, userID, planID uint64,
	status string,
	startDate time.Time,
	endDate *time.Time,
	autoRenew bool,
	createdAt, updatedAt time.Time,
) (SubscriptionModel, error) {
	if err := validateSubscription(userID, planID, startDate, endDate); err != nil {
		return SubscriptionModel{}, err
	}

	statusEnum, err := enum.NewSubscriptionStatusEnum(status)
	if err != nil {
		return SubscriptionModel{}, err
	}

	return SubscriptionModel{
		id:        id,
		userID:    userID,
		planID:    planID,
		status:    statusEnum,
		startDate: startDate,
		endDate:   endDate,
		autoRenew: autoRenew,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (s *SubscriptionModel) ID() uint64 {
	return s.id
}

func (s *SubscriptionModel) UserID() uint64 {
	return s.userID
}

func (s *SubscriptionModel) PlanID() uint64 {
	return s.planID
}

func (s *SubscriptionModel) Status() enum.SubscriptionStatusEnum {
	return s.status
}

func (s *SubscriptionModel) StartDate() time.Time {
	return s.startDate
}

func (s *SubscriptionModel) EndDate() *time.Time {
	return s.endDate
}

func (s *SubscriptionModel) AutoRenew() bool {
	return s.autoRenew
}

func (s *SubscriptionModel) CreatedAt() time.Time {
	return s.createdAt
}

func (s *SubscriptionModel) UpdatedAt() time.Time {
	return s.updatedAt
}

func (s *SubscriptionModel) UpdateStatus(statusValue string) error {
	status, err := enum.NewSubscriptionStatusEnum(statusValue)
	if err != nil {
		return err
	}

	s.status = status
	s.updatedAt = time.Now().UTC()
	return nil
}

func (s *SubscriptionModel) UpdateEndDate(endDate *time.Time) error {
	if endDate != nil && endDate.Before(s.startDate) {
		return errs.ErrEndDateBeforeStartDate
	}

	s.endDate = endDate
	s.updatedAt = time.Now().UTC()
	return nil
}

func (s *SubscriptionModel) SetAutoRenew(autoRenew bool) {
	s.autoRenew = autoRenew
	s.updatedAt = time.Now().UTC()
}

func validateSubscription(
	userID, planID uint64,
	startDate time.Time,
	endDate *time.Time,
) error {
	if userID == 0 {
		return errs.ErrUserIDRequired
	}

	if planID == 0 {
		return errs.ErrPlanIDRequired
	}

	if startDate.IsZero() {
		return errs.ErrStartDateRequired
	}

	if endDate != nil && !startDate.IsZero() && endDate.Before(startDate) {
		return errs.ErrEndDateBeforeStartDate
	}

	return nil
}
