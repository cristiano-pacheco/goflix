package model

import (
	"strings"
	"time"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/enum"
)

type PlanModel struct {
	id          uint64
	name        NameModel
	description *DescriptionModel
	amount      AmountModel
	currency    CurrencyModel
	interval    enum.PlanIntervalEnum
	trialPeriod *TrialPeriodModel
	createdAt   time.Time
	updatedAt   time.Time
}

func CreatePlanModel(
	name, description, currency, interval string,
	amountCents uint,
	trialPeriod *uint,
) (PlanModel, error) {
	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)
	currency = strings.TrimSpace(currency)

	nameModel, err := CreateNameModel(name)
	if err != nil {
		return PlanModel{}, err
	}

	var descriptionModel *DescriptionModel
	if description != "" {
		desc, err := CreateDescriptionModel(description)
		if err != nil {
			return PlanModel{}, err
		}
		descriptionModel = &desc
	}

	amountModel, err := CreateAmountModel(amountCents)
	if err != nil {
		return PlanModel{}, err
	}

	currencyModel, err := CreateCurrencyModel(currency)
	if err != nil {
		return PlanModel{}, err
	}

	var trialPeriodModel *TrialPeriodModel
	if trialPeriod != nil {
		trial, err := CreateTrialPeriodModel(*trialPeriod)
		if err != nil {
			return PlanModel{}, err
		}
		trialPeriodModel = &trial
	}

	planInterval, err := enum.NewPlanIntervalEnum(interval)
	if err != nil {
		return PlanModel{}, err
	}

	return PlanModel{
		name:        nameModel,
		description: descriptionModel,
		amount:      amountModel,
		currency:    currencyModel,
		interval:    planInterval,
		trialPeriod: trialPeriodModel,
		createdAt:   time.Now().UTC(),
		updatedAt:   time.Now().UTC(),
	}, nil
}

func RestorePlanModel(
	id uint64,
	name, description, currency, interval string,
	amountCents uint,
	trialPeriod *uint,
	createdAt, updatedAt time.Time,
) (PlanModel, error) {
	nameModel, err := CreateNameModel(name)
	if err != nil {
		return PlanModel{}, err
	}

	var descriptionModel *DescriptionModel
	if description != "" {
		desc, err := CreateDescriptionModel(description)
		if err != nil {
			return PlanModel{}, err
		}
		descriptionModel = &desc
	}

	amountModel, err := CreateAmountModel(amountCents)
	if err != nil {
		return PlanModel{}, err
	}

	currencyModel, err := CreateCurrencyModel(currency)
	if err != nil {
		return PlanModel{}, err
	}

	var trialPeriodModel *TrialPeriodModel
	if trialPeriod != nil {
		trial, err := CreateTrialPeriodModel(*trialPeriod)
		if err != nil {
			return PlanModel{}, err
		}
		trialPeriodModel = &trial
	}

	planInterval, err := enum.NewPlanIntervalEnum(interval)
	if err != nil {
		return PlanModel{}, err
	}

	return PlanModel{
		id:          id,
		name:        nameModel,
		description: descriptionModel,
		amount:      amountModel,
		currency:    currencyModel,
		interval:    planInterval,
		trialPeriod: trialPeriodModel,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil
}

func (p PlanModel) ID() uint64 {
	return p.id
}

func (p PlanModel) Name() NameModel {
	return p.name
}

func (p PlanModel) Description() *DescriptionModel {
	return p.description
}

func (p PlanModel) Amount() AmountModel {
	return p.amount
}

func (p PlanModel) Currency() CurrencyModel {
	return p.currency
}

func (p PlanModel) Interval() enum.PlanIntervalEnum {
	return p.interval
}

func (p PlanModel) TrialPeriod() *TrialPeriodModel {
	return p.trialPeriod
}

func (p PlanModel) CreatedAt() time.Time {
	return p.createdAt
}

func (p PlanModel) UpdatedAt() time.Time {
	return p.updatedAt
}
