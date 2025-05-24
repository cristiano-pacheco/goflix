package repository

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
)

type AuthTokenRepository interface {
	Create(ctx context.Context, authToken model.AuthTokenModel) (model.AuthTokenModel, error)
	Update(ctx context.Context, authToken model.AuthTokenModel) error
	Delete(ctx context.Context, id uint64) error
	FindByToken(ctx context.Context, token string) (model.AuthTokenModel, error)
}
