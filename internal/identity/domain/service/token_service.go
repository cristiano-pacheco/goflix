package service

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
)

type TokenService interface {
	Generate(ctx context.Context, user model.UserModel) (string, error)
}
