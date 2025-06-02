package model

import (
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
)

const (
	minTrialPeriodDays = 1
	maxTrialPeriodDays = 30
)

type TrialPeriodModel struct {
	value uint
}

func CreateTrialPeriodModel(value uint) (TrialPeriodModel, error) {
	if err := validateTrialPeriod(value); err != nil {
		return TrialPeriodModel{}, err
	}
	return TrialPeriodModel{value: value}, nil
}

func (t *TrialPeriodModel) Days() uint {
	return t.value
}

func validateTrialPeriod(value uint) error {
	if value < minTrialPeriodDays {
		return errs.ErrTrialPeriodTooShort
	}

	if value > maxTrialPeriodDays {
		return errs.ErrTrialPeriodTooLong
	}

	return nil
}
