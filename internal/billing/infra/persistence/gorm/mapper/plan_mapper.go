package mapper

import (
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/entity"
)

type PlanMapper interface {
	ToModel(entity entity.PlanEntity) (model.PlanModel, error)
	ToEntity(model model.PlanModel) entity.PlanEntity
}

type planMapper struct {
}

func NewPlanMapper() PlanMapper {
	return &planMapper{}
}

func (p *planMapper) ToModel(entity entity.PlanEntity) (model.PlanModel, error) {
	var trialPeriod *uint
	if entity.TrialPeriod > 0 {
		trialPeriod = &entity.TrialPeriod
	}

	planModel, err := model.RestorePlanModel(
		entity.ID,
		entity.Name,
		entity.Description,
		entity.Currency,
		entity.Interval,
		entity.AmountCents,
		trialPeriod,
		entity.CreatedAt,
		entity.UpdatedAt,
	)
	if err != nil {
		return model.PlanModel{}, err
	}
	return planModel, nil
}

func (p *planMapper) ToEntity(model model.PlanModel) entity.PlanEntity {
	var trialPeriod uint
	if model.TrialPeriod() != nil {
		trialPeriod = model.TrialPeriod().Days()
	}

	var description string
	if model.Description() != nil {
		description = model.Description().String()
	}

	nameModel := model.Name()
	currencyModel := model.Currency()
	amountModel := model.Amount()
	intervalModel := model.Interval()

	return entity.PlanEntity{
		ID:          model.ID(),
		Name:        (&nameModel).String(),
		Description: description,
		AmountCents: (&amountModel).Cents(),
		Currency:    currencyModel.Code(),
		Interval:    intervalModel.String(),
		TrialPeriod: trialPeriod,
		CreatedAt:   model.CreatedAt(),
		UpdatedAt:   model.UpdatedAt(),
	}
}
