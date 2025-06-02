package model

import (
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
)

const (
	maxAmountCents = 999999999 // 9,999,999.99 in major currency units
)

type AmountModel struct {
	value uint
}

func CreateAmountModel(value uint) (AmountModel, error) {
	if err := validateAmount(value); err != nil {
		return AmountModel{}, err
	}
	return AmountModel{value: value}, nil
}

func (a *AmountModel) Cents() uint {
	return a.value
}

func validateAmount(value uint) error {
	if value > maxAmountCents {
		return errs.ErrAmountExceedsMaximum
	}

	return nil
}
