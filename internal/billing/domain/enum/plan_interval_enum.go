package enum

import "errors"

// Constants with pattern: Enum{EnumName}{Value}
const (
	EnumPlanIntervalDay   string = "Day"
	EnumPlanIntervalWeek  string = "Week"
	EnumPlanIntervalMonth string = "Month"
	EnumPlanIntervalYear  string = "Year"
)

// Struct type: {EnumName}Enum
type PlanIntervalEnum struct {
	value string
}

// Constructor: New{EnumName}Enum
func NewPlanIntervalEnum(value string) (PlanIntervalEnum, error) {
	if err := validatePlanIntervalEnum(value); err != nil {
		return PlanIntervalEnum{}, err
	}

	return PlanIntervalEnum{value: value}, nil
}

// String method
func (e *PlanIntervalEnum) String() string {
	return e.value
}

// Validation function
func validatePlanIntervalEnum(value string) error {
	allowedValues := map[string]struct{}{
		EnumPlanIntervalDay:   {},
		EnumPlanIntervalWeek:  {},
		EnumPlanIntervalMonth: {},
		EnumPlanIntervalYear:  {},
	}

	if _, ok := allowedValues[value]; !ok {
		return errors.New("invalid plan interval: " + value)
	}

	return nil
}
