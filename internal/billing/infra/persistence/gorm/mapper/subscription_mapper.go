package mapper

import (
	"time"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/entity"
)

type SubscriptionMapper interface {
	ToModel(entity entity.SubscriptionEntity) (model.SubscriptionModel, error)
	ToEntity(model model.SubscriptionModel) entity.SubscriptionEntity
}

type subscriptionMapper struct {
}

func NewSubscriptionMapper() SubscriptionMapper {
	return &subscriptionMapper{}
}

func (s *subscriptionMapper) ToModel(entity entity.SubscriptionEntity) (model.SubscriptionModel, error) {
	var endDate *time.Time
	if !entity.EndDate.IsZero() {
		endDate = &entity.EndDate
	}

	subscriptionModel, err := model.RestoreSubscriptionModel(
		entity.ID,
		entity.UserID,
		entity.PlanID,
		entity.Status,
		entity.StartDate,
		endDate,
		entity.AutoRenew,
		entity.CreatedAt,
		entity.UpdatedAt,
	)
	if err != nil {
		return model.SubscriptionModel{}, err
	}
	return subscriptionModel, nil
}

func (s *subscriptionMapper) ToEntity(model model.SubscriptionModel) entity.SubscriptionEntity {
	var endDate time.Time
	if model.EndDate() != nil {
		endDate = *model.EndDate()
	}

	statusEnum := model.Status()

	return entity.SubscriptionEntity{
		ID:        model.ID(),
		UserID:    model.UserID(),
		PlanID:    model.PlanID(),
		Status:    (&statusEnum).String(),
		StartDate: model.StartDate(),
		EndDate:   endDate,
		AutoRenew: model.AutoRenew(),
		CreatedAt: model.CreatedAt(),
		UpdatedAt: model.UpdatedAt(),
	}
}
