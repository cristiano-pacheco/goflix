package repository

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.UserModel) (model.UserModel, error)
	Update(ctx context.Context, user model.UserModel) error

	FindByEmail(ctx context.Context, email string) (model.UserModel, error)
	FindByID(ctx context.Context, id uint64) (model.UserModel, error)
	FindByConfirmationToken(ctx context.Context, token string) (model.UserModel, error)
	FindByResetPasswordToken(ctx context.Context, token string) (model.UserModel, error)
	IsActivated(ctx context.Context, id uint64) (bool, error)
}
