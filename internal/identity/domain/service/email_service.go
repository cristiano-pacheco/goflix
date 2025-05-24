package service

import "github.com/cristiano-pacheco/goflix/internal/identity/domain/model"

type EmailService interface {
	SendAccountConfirmationEmail(user model.UserModel, token string) error
}
