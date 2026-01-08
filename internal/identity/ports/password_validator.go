package ports

type PasswordValidatorI interface {
	Validate(password string) error
}
