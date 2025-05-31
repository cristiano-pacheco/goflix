package repository

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
)

type PlanRepository interface {
	Create(ctx context.Context, plan model.PlanModel) (model.PlanModel, error)
	Update(ctx context.Context, plan model.PlanModel) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (model.PlanModel, error)
	FindAll(ctx context.Context) ([]model.PlanModel, error)
}
