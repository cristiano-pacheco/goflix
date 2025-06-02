package enum

import (
	"fmt"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
)

// Constants with pattern: Enum{EnumName}{Value}
const (
	EnumPlanIntervalDay   string = "Day"
	EnumPlanIntervalWeek  string = "Week"
	EnumPlanIntervalMonth string = "Month"
	EnumPlanIntervalYear  string = "Year"
)

type PlanIntervalEnum struct {
	value string
}

func NewPlanIntervalEnum(value string) (PlanIntervalEnum, error) {
	if err := validatePlanIntervalEnum(value); err != nil {
		return PlanIntervalEnum{}, err
	}

	return PlanIntervalEnum{value: value}, nil
}

func (e *PlanIntervalEnum) String() string {
	return e.value
}

func validatePlanIntervalEnum(value string) error {
	allowedValues := map[string]struct{}{
		EnumPlanIntervalDay:   {},
		EnumPlanIntervalWeek:  {},
		EnumPlanIntervalMonth: {},
		EnumPlanIntervalYear:  {},
	}

	if _, ok := allowedValues[value]; !ok {
		return fmt.Errorf("%w: %s", errs.ErrInvalidPlanInterval, value)
	}

	return nil
}
