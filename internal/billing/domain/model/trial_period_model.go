package model

import (
	"errors"
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
		return errors.New("trial period must be at least 1 day")
	}

	if value > maxTrialPeriodDays {
		return errors.New("trial period cannot exceed 365 days")
	}

	return nil
}
