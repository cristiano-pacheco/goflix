package usecase

import (
	"context"
	"errors"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/service"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/validator"
)

type TokenGenerateUseCase struct {
	validator    validator.Validate
	userRepo     repository.UserRepository
	hashService  service.HashService
	tokenService service.TokenService
}

func NewTokenGenerateUseCase(
	validator validator.Validate,
	userRepo repository.UserRepository,
	hashService service.HashService,
	tokenService service.TokenService,
) *TokenGenerateUseCase {
	return &TokenGenerateUseCase{
		validator,
		userRepo,
		hashService,
		tokenService,
	}
}

type TokenGenerateInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type TokenGenerateOutput struct {
	Token string
}

func (uc *TokenGenerateUseCase) Execute(ctx context.Context, input TokenGenerateInput) (TokenGenerateOutput, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "TokenGenerateUseCase.Execute")
	defer span.End()

	output := TokenGenerateOutput{}

	err := uc.validator.Struct(input)
	if err != nil {
		return output, err
	}

	user, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return output, errs.ErrInvalidCredentials
		}
		return output, err
	}

	if !user.IsActivated() {
		return output, errs.ErrUserIsNotActivated
	}

	hash := []byte(user.PasswordHash())
	pass := []byte(input.Password)
	err = uc.hashService.CompareHashAndPassword(hash, pass)
	if err != nil {
		return output, errs.ErrInvalidCredentials
	}

	token, err := uc.tokenService.Generate(ctx, user)
	if err != nil {
		return output, err
	}

	output.Token = token
	return output, nil
}
