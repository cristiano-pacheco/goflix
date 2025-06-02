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

type PlanRepository interface {
	repository.PlanRepository
}

type planRepository struct {
	db     *database.GoflixDB
	mapper mapper.PlanMapper
}

func NewPlanRepository(db *database.GoflixDB, mapper mapper.PlanMapper) PlanRepository {
	return &planRepository{db, mapper}
}

func (r *planRepository) Create(ctx context.Context, planModel model.PlanModel) (model.PlanModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "PlanRepository.Create")
	defer span.End()

	planEntity := r.mapper.ToEntity(planModel)
	result := r.db.WithContext(ctx).Create(&planEntity)
	if result.Error != nil {
		return model.PlanModel{}, result.Error
	}

	planModel, err := r.mapper.ToModel(planEntity)
	if err != nil {
		return model.PlanModel{}, err
	}

	return planModel, nil
}

func (r *planRepository) Update(ctx context.Context, planModel model.PlanModel) error {
	ctx, span := otel.Trace().StartSpan(ctx, "PlanRepository.Update")
	defer span.End()

	planEntity := r.mapper.ToEntity(planModel)
	result := r.db.WithContext(ctx).Save(&planEntity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *planRepository) Delete(ctx context.Context, id uint64) error {
	ctx, span := otel.Trace().StartSpan(ctx, "PlanRepository.Delete")
	defer span.End()

	result := r.db.WithContext(ctx).Delete(&entity.PlanEntity{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *planRepository) FindByID(ctx context.Context, id uint64) (model.PlanModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "PlanRepository.FindByID")
	defer span.End()

	var planEntity entity.PlanEntity
	r.db.WithContext(ctx).First(&planEntity, id)
	if planEntity.ID == 0 {
		return model.PlanModel{}, errs.ErrNotFound
	}

	planModel, err := r.mapper.ToModel(planEntity)
	if err != nil {
		return model.PlanModel{}, err
	}

	return planModel, nil
}

func (r *planRepository) FindAll(ctx context.Context) ([]model.PlanModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "PlanRepository.FindAll")
	defer span.End()

	var planEntities []entity.PlanEntity
	result := r.db.WithContext(ctx).Find(&planEntities)
	if result.Error != nil {
		return nil, result.Error
	}

	planModels := make([]model.PlanModel, 0, len(planEntities))
	for _, planEntity := range planEntities {
		planModel, err := r.mapper.ToModel(planEntity)
		if err != nil {
			return nil, err
		}
		planModels = append(planModels, planModel)
	}

	return planModels, nil
}
