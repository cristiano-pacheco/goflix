package identity

import (
	"go.uber.org/fx"

	"github.com/cristiano-pacheco/goflix/internal/identity/application/usecase"
	domain_service "github.com/cristiano-pacheco/goflix/internal/identity/domain/service"
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/validator"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/handler"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/router"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/repository"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/service"
	"github.com/cristiano-pacheco/goflix/internal/identity/ports"
)

var Module = fx.Module(
	"identity",
	fx.Provide(
		usecase.NewUserCreateUseCase,
		usecase.NewUserActivateUseCase,
		usecase.NewUserUpdateUseCase,
		usecase.NewUserFindUseCase,
		usecase.NewTokenGenerateUseCase,
		handler.NewAuthHandler,
		handler.NewUserHandler,
		fx.Annotate(
			domain_service.NewHashService,
			fx.As(new(domain_service.HashServiceI)),
		),
		fx.Annotate(
			validator.NewPasswordValidator,
			fx.As(new(ports.PasswordValidatorI)),
		),
		fx.Annotate(
			mapper.NewAuthTokenMapper,
			fx.As(new(mapper.AuthTokenMapperI)),
		),
		fx.Annotate(
			mapper.NewUserMapper,
			fx.As(new(mapper.UserMapperI)),
		),
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(ports.UserRepositoryI)),
		),
		fx.Annotate(
			repository.NewAuthTokenRepository,
			fx.As(new(ports.AuthTokenRepositoryI)),
		),
		fx.Annotate(
			service.NewSendEmailConfirmationService,
			fx.As(new(ports.SendEmailConfirmationServiceI)),
		),
		fx.Annotate(
			service.NewTokenService,
			fx.As(new(ports.TokenServiceI)),
		),
	),
	fx.Invoke(
		router.SetupUserRoutes,
		router.SetupAuthRoutes,
	),
)
