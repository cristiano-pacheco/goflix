package usecase

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/logger"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
)

type UserFindUseCase struct {
	userRepo repository.UserRepository
	logger   logger.Logger
}

func NewUserFindUseCase(
	userRepo repository.UserRepository,
	logger logger.Logger,
) *UserFindUseCase {
	return &UserFindUseCase{userRepo, logger}
}

type UserFindInput struct {
	UserID uint64 `validate:"required,number"`
}

type UserFindOutput struct {
	UserID   uint64
	Name     string
	Email    string
	Password string
}

func (uc *UserFindUseCase) Execute(ctx context.Context, input UserFindInput) (UserFindOutput, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserFindUseCase.Execute")
	defer span.End()

	userModel, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		message := "error finding user by id %d"
		uc.logger.Error(message, "error", err, "id", input.UserID)
		return UserFindOutput{}, err
	}

	output := UserFindOutput{
		UserID:   userModel.ID(),
		Name:     userModel.Name(),
		Email:    userModel.Email(),
		Password: userModel.PasswordHash(),
	}

	return output, nil
}
