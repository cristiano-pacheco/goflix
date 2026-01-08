package ports

import "github.com/cristiano-pacheco/goflix/internal/identity/domain/model"

type EmailServiceI interface {
	SendAccountConfirmationEmail(user model.UserModel, token string) error
}
