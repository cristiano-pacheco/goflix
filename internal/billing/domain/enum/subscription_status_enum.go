package enum

import "errors"

const (
	EnumSubscriptionStatusActive    string = "Active"
	EnumSubscriptionStatusInactive  string = "Inactive"
	EnumSubscriptionStatusCancelled string = "Cancelled"
	EnumSubscriptionStatusExpired   string = "Expired"
	EnumSubscriptionStatusPastDue   string = "PastDue"
)

type SubscriptionStatusEnum struct {
	value string
}

func NewSubscriptionStatusEnum(value string) (SubscriptionStatusEnum, error) {
	if err := validateSubscriptionStatusEnum(value); err != nil {
		panic(err)
	}

	return SubscriptionStatusEnum{value: value}, nil
}

func (s *SubscriptionStatusEnum) String() string {
	return s.value
}

func validateSubscriptionStatusEnum(value string) error {
	allowedValues := map[string]struct{}{
		EnumSubscriptionStatusActive:    {},
		EnumSubscriptionStatusInactive:  {},
		EnumSubscriptionStatusCancelled: {},
		EnumSubscriptionStatusExpired:   {},
		EnumSubscriptionStatusPastDue:   {},
	}

	if _, ok := allowedValues[value]; !ok {
		return errors.New("invalid subscription status: " + value)
	}

	return nil
}
