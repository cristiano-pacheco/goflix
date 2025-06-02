package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/enum"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/mapper"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/repository"
	sharedErrs "github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/logger"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/validator"
)

type CreateSubscriptionUseCase struct {
	subscriptionRepository repository.SubscriptionRepository
	planRepository         repository.PlanRepository
	endDateMapper          mapper.EndDateMapper
	validate               validator.Validate
	logger                 logger.Logger
}

func NewCreateSubscriptionUseCase(
	subscriptionRepository repository.SubscriptionRepository,
	planRepository repository.PlanRepository,
	endDateMapper mapper.EndDateMapper,
	validate validator.Validate,
	logger logger.Logger,
) *CreateSubscriptionUseCase {
	return &CreateSubscriptionUseCase{
		subscriptionRepository,
		planRepository,
		endDateMapper,
		validate,
		logger,
	}
}

type CreateSubscriptionInput struct {
	PlanID uint64 `validate:"required,number"`
	UserID uint64 `validate:"required,number"`
}

type CreateSubscriptionOutput struct {
	SubscriptionID uint64
	UserID         uint64
	PlanID         uint64
	Status         string
	StartDate      time.Time
	EndDate        *time.Time
	AutoRenew      bool
}

func (uc *CreateSubscriptionUseCase) Execute(
	ctx context.Context,
	input CreateSubscriptionInput,
) (CreateSubscriptionOutput, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "CreateSubscriptionUseCase.Execute")
	defer span.End()

	output := CreateSubscriptionOutput{}

	err := uc.validate.Struct(input)
	if err != nil {
		return output, err
	}

	// Verify the plan exists
	planModel, err := uc.planRepository.FindByID(ctx, input.PlanID)
	if err != nil {
		if errors.Is(err, errs.ErrPlanNotFound) {
			message := "plan not found with id %d"
			uc.logger.Error(message, "error", err, "planID", input.PlanID)
			return output, errs.ErrPlanNotFound
		}
		message := "error finding plan by id %d"
		uc.logger.Error(message, "error", err, "planID", input.PlanID)
		return output, err
	}

	// Check if user already has an active subscription
	existingSubscriptions, err := uc.subscriptionRepository.FindByUserID(ctx, input.UserID)
	if err != nil && !errors.Is(err, sharedErrs.ErrNotFound) {
		message := "error finding existing subscriptions for user %d"
		uc.logger.Error(message, "error", err, "userID", input.UserID)
		return output, err
	}

	// Check for active subscriptions
	for _, subscription := range existingSubscriptions {
		status := subscription.Status()
		if status.String() == enum.EnumSubscriptionStatusActive {
			return output, errs.ErrUserAlreadyHasActiveSubscription
		}
	}

	// Calculate subscription dates based on plan interval
	startDate := time.Now().UTC()
	endDate := uc.endDateMapper.Map(startDate, planModel.Interval())

	// Create subscription model
	subscriptionModel, err := model.CreateSubscriptionModel(
		input.UserID,
		input.PlanID,
		startDate,
		endDate,
	)
	if err != nil {
		message := "error creating subscription model"
		uc.logger.Error(message, "error", err)
		return output, err
	}

	// TODO: charge user

	// Save subscription
	createdSubscription, err := uc.subscriptionRepository.Create(ctx, subscriptionModel)
	if err != nil {
		message := "error creating subscription"
		uc.logger.Error(message, "error", err)
		return output, err
	}

	// TODO: send email to user

	createdStatus := createdSubscription.Status()
	output = CreateSubscriptionOutput{
		SubscriptionID: createdSubscription.ID(),
		UserID:         createdSubscription.UserID(),
		PlanID:         createdSubscription.PlanID(),
		Status:         createdStatus.String(),
		StartDate:      createdSubscription.StartDate(),
		EndDate:        createdSubscription.EndDate(),
		AutoRenew:      createdSubscription.AutoRenew(),
	}

	return output, nil
}
