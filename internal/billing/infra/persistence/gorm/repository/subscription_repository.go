package repository

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/repository"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/entity"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/database"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
)

type SubscriptionRepository interface {
	repository.SubscriptionRepository
}

type subscriptionRepository struct {
	db     *database.GoflixDB
	mapper mapper.SubscriptionMapper
}

func NewSubscriptionRepository(db *database.GoflixDB, mapper mapper.SubscriptionMapper) SubscriptionRepository {
	return &subscriptionRepository{db, mapper}
}

func (r *subscriptionRepository) Create(ctx context.Context, subscriptionModel model.SubscriptionModel) (model.SubscriptionModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "SubscriptionRepository.Create")
	defer span.End()

	subscriptionEntity := r.mapper.ToEntity(subscriptionModel)
	result := r.db.WithContext(ctx).Create(&subscriptionEntity)
	if result.Error != nil {
		return model.SubscriptionModel{}, result.Error
	}

	subscriptionModel, err := r.mapper.ToModel(subscriptionEntity)
	if err != nil {
		return model.SubscriptionModel{}, err
	}

	return subscriptionModel, nil
}

func (r *subscriptionRepository) Update(ctx context.Context, subscriptionModel model.SubscriptionModel) error {
	ctx, span := otel.Trace().StartSpan(ctx, "SubscriptionRepository.Update")
	defer span.End()

	subscriptionEntity := r.mapper.ToEntity(subscriptionModel)
	result := r.db.WithContext(ctx).Save(&subscriptionEntity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *subscriptionRepository) Delete(ctx context.Context, id uint64) error {
	ctx, span := otel.Trace().StartSpan(ctx, "SubscriptionRepository.Delete")
	defer span.End()

	result := r.db.WithContext(ctx).Delete(&entity.SubscriptionEntity{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *subscriptionRepository) FindByID(ctx context.Context, id uint64) (model.SubscriptionModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "SubscriptionRepository.FindByID")
	defer span.End()

	var subscriptionEntity entity.SubscriptionEntity
	r.db.WithContext(ctx).First(&subscriptionEntity, id)
	if subscriptionEntity.ID == 0 {
		return model.SubscriptionModel{}, errs.ErrNotFound
	}

	subscriptionModel, err := r.mapper.ToModel(subscriptionEntity)
	if err != nil {
		return model.SubscriptionModel{}, err
	}

	return subscriptionModel, nil
}

func (r *subscriptionRepository) FindByUserID(ctx context.Context, userID uint64) ([]model.SubscriptionModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "SubscriptionRepository.FindByUserID")
	defer span.End()

	var subscriptionEntities []entity.SubscriptionEntity
	result := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&subscriptionEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	subscriptionModels := make([]model.SubscriptionModel, 0, len(subscriptionEntities))
	for _, subscriptionEntity := range subscriptionEntities {
		subscriptionModel, err := r.mapper.ToModel(subscriptionEntity)
		if err != nil {
			return nil, err
		}
		subscriptionModels = append(subscriptionModels, subscriptionModel)
	}

	return subscriptionModels, nil
}

func (r *subscriptionRepository) FindActiveSubscriptionByUserID(ctx context.Context, userID uint64) (model.SubscriptionModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "SubscriptionRepository.FindActiveSubscriptionByUserID")
	defer span.End()

	var subscriptionEntity entity.SubscriptionEntity
	result := r.db.WithContext(ctx).Where("user_id = ? AND status = ?", userID, "Active").First(&subscriptionEntity)
	if result.Error != nil {
		return model.SubscriptionModel{}, result.Error
	}

	if subscriptionEntity.ID == 0 {
		return model.SubscriptionModel{}, errs.ErrNotFound
	}

	subscriptionModel, err := r.mapper.ToModel(subscriptionEntity)
	if err != nil {
		return model.SubscriptionModel{}, err
	}

	return subscriptionModel, nil
} 