package identity

import (
	"go.uber.org/fx"

	"github.com/cristiano-pacheco/goflix/internal/identity/application/usecase"
	domain_repository "github.com/cristiano-pacheco/goflix/internal/identity/domain/repository"
	domain_service "github.com/cristiano-pacheco/goflix/internal/identity/domain/service"
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/validator"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/handler"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/middleware"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/router"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/repository"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/service"
)

var Module = fx.Module(
	"identity",
	fx.Provide(
		// #################### APPLICATION ####################################
		// usecases
		usecase.NewUserCreateUseCase,
		usecase.NewUserActivateUseCase,
		usecase.NewUserUpdateUseCase,
		usecase.NewUserFindUseCase,
		usecase.NewTokenGenerateUseCase,

		// #################### DOMAIN #########################################
		domain_service.NewHashService,
		validator.NewPasswordValidator,

		// #################### INFRA ##########################################
		router.NewRouter,

		// handlers
		handler.NewAuthHandler,
		handler.NewUserHandler,

		// middlewares
		middleware.NewAuthMiddleware,

		// mappers
		mapper.NewUserMapper,
		mapper.NewAuthTokenMapper,

		// repositories
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(domain_repository.UserRepository)),
		),

		fx.Annotate(
			repository.NewAuthTokenRepository,
			fx.As(new(domain_repository.AuthTokenRepository)),
		),

		// services
		fx.Annotate(
			service.NewSendEmailConfirmationService,
			fx.As(new(domain_service.SendEmailConfirmationService)),
		),

		fx.Annotate(
			service.NewTokenService,
			fx.As(new(domain_service.TokenService)),
		),
	),
	fx.Invoke(
		router.SetupUserRoutes,
		router.SetupAuthRoutes,
	),
)
