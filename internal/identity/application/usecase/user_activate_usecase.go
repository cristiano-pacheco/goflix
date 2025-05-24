package usecase

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/logger"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/validator"
)

type UserActivateUseCase struct {
	userRepository repository.UserRepository
	validate       validator.Validate
	logger         logger.Logger
}

func NewUserActivateUseCase(
	userRepository repository.UserRepository,
	validate validator.Validate,
	logger logger.Logger,
) *UserActivateUseCase {
	return &UserActivateUseCase{userRepository, validate, logger}
}

type UserActivateInput struct {
	Token string `validate:"required"`
}

func (uc *UserActivateUseCase) Execute(ctx context.Context, input UserActivateInput) error {
	ctx, span := otel.Trace().StartSpan(ctx, "UserActivateUseCase.Execute")
	defer span.End()

	err := uc.validate.Struct(input)
	if err != nil {
		return err
	}

	user, err := uc.userRepository.FindByConfirmationToken(ctx, input.Token)
	if err != nil {
		return err
	}

	if !user.IsConfirmationTokenValid(input.Token) {
		return errs.ErrInvalidAccountConfirmationToken
	}

	user.ConfirmAccount()
	err = uc.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
