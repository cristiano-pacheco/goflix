package usecase

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/service"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/logger"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/validator"
)

type UserUpdateUseCase struct {
	validate    validator.Validate
	userRepo    repository.UserRepository
	logger      logger.Logger
	hashService service.HashService
}

func NewUserUpdateUseCase(
	validate validator.Validate,
	userRepo repository.UserRepository,
	logger logger.Logger,
	hashService service.HashService,
) *UserUpdateUseCase {
	return &UserUpdateUseCase{validate, userRepo, logger, hashService}
}

type UserUpdateInput struct {
	UserID   uint64 `validate:"required"`
	Name     string `validate:"required,min=3,max=255"`
	Password string `validate:"required,min=8"`
}

func (uc *UserUpdateUseCase) Execute(ctx context.Context, input UserUpdateInput) error {
	ctx, span := otel.Trace().StartSpan(ctx, "UserUpdateUseCase.Execute")
	defer span.End()

	err := uc.validate.Struct(input)
	if err != nil {
		return err
	}

	userModel, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		return err
	}

	ph, err := uc.hashService.GenerateFromPassword([]byte(input.Password))
	if err != nil {
		message := "error generating password hash"
		uc.logger.Error(message, "error", err)
		return err
	}

	updatedUserModel, err := model.RestoreUserModel(
		userModel.ID(),
		input.Name,
		userModel.Email(),
		string(ph),
		userModel.IsActivated(),
		userModel.ResetPasswordToken(),
		userModel.ResetPasswordExpiresAt(),
		userModel.ConfirmedAt(),
		userModel.ConfirmationToken(),
		userModel.ConfirmationExpiresAt(),
		userModel.CreatedAt(),
		userModel.UpdatedAt(),
	)
	if err != nil {
		return err
	}

	err = uc.userRepo.Update(ctx, updatedUserModel)
	if err != nil {
		message := "error updating user with id %d"
		uc.logger.Error(message, "error", err, "id", input.UserID)
		return err
	}

	return nil
}
