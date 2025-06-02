package errs

import "errors"

var (
	ErrPlanNotFound = errors.New("plan not found")

	ErrAmountExceedsMaximum = errors.New("amount exceeds maximum allowed value")

	ErrNameRequired                   = errors.New("plan name is required")
	ErrNameTooShort                   = errors.New("plan name must be at least 2 characters long")
	ErrNameTooLong                    = errors.New("plan name cannot exceed 100 characters")
	ErrNameMustStartWithLetterOrDigit = errors.New("plan name must start with a letter or digit")
	ErrNameMustEndWithLetterOrDigit   = errors.New("plan name must end with a letter or digit")
	ErrNameConsecutiveSpaces          = errors.New("plan name cannot contain consecutive spaces")
	ErrNameInvalidCharacters          = errors.New(
		"plan name contains invalid characters (only letters, digits, spaces, hyphens, underscores, and periods are allowed)",
	)
	ErrNameCannotStartOrEndWithSpaces      = errors.New("plan name cannot start or end with spaces")
	ErrNameCannotStartWithPunctuation      = errors.New("plan name cannot start with punctuation")
	ErrNameCannotEndWithPunctuation        = errors.New("plan name cannot end with punctuation")
	ErrNameExcessiveConsecutivePunctuation = errors.New(
		"plan name cannot contain more than 2 consecutive punctuation marks",
	)

	ErrDescriptionTooLong           = errors.New("description cannot exceed 255 characters")
	ErrDescriptionInvalidCharacters = errors.New(
		"description contains invalid characters (only printable characters are allowed)",
	)
	ErrDescriptionCannotStartOrEndWithSpaces = errors.New("description cannot start or end with spaces")
	ErrDescriptionExcessiveConsecutiveSpaces = errors.New("description cannot contain more than 2 consecutive spaces")
	ErrDescriptionControlCharacters          = errors.New("description cannot contain control characters")

	ErrCurrencyCodeEmpty         = errors.New("currency code cannot be empty")
	ErrCurrencyCodeInvalidLength = errors.New("currency code must be exactly 3 characters")
	ErrCurrencyCodeInvalid       = errors.New("invalid currency code")

	ErrTrialPeriodTooShort = errors.New("trial period must be at least 1 day")
	ErrTrialPeriodTooLong  = errors.New("trial period cannot exceed 365 days")

	ErrSubscriptionNotFound = errors.New("subscription not found")

	ErrUserIDRequired         = errors.New("user ID is required")
	ErrPlanIDRequired         = errors.New("plan ID is required")
	ErrStartDateRequired      = errors.New("start date is required")
	ErrEndDateBeforeStartDate = errors.New("end date cannot be before start date")

	ErrInvalidSubscriptionStatus        = errors.New("invalid subscription status")
	ErrInvalidPlanInterval              = errors.New("invalid plan interval")
	ErrUserAlreadyHasActiveSubscription = errors.New("user already has an active subscription")
)
