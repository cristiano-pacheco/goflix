package repository

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
)

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription model.SubscriptionModel) (model.SubscriptionModel, error)
	Update(ctx context.Context, subscription model.SubscriptionModel) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (model.SubscriptionModel, error)
	FindByUserID(ctx context.Context, userID uint64) ([]model.SubscriptionModel, error)
}
