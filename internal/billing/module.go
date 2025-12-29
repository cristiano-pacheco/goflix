package billing

import (
	"go.uber.org/fx"

	"github.com/cristiano-pacheco/goflix/internal/billing/application/usecase"
	domain_mapper "github.com/cristiano-pacheco/goflix/internal/billing/domain/mapper"
	domain_repository "github.com/cristiano-pacheco/goflix/internal/billing/domain/repository"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/http/handler"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/http/router"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/repository"
)

var Module = fx.Module(
	"billing",
	fx.Provide(
		usecase.NewSubscriptionCreateUseCase,
		handler.NewSubscriptionHandler,
		fx.Annotate(
			NewBillingFacade,
			fx.As(new(BillingFacadeI)),
		),
		fx.Annotate(
			domain_mapper.NewEndDateMapper,
			fx.As(new(domain_mapper.EndDateMapperI)),
		),
		fx.Annotate(
			mapper.NewSubscriptionMapper,
			fx.As(new(mapper.SubscriptionMapperI)),
		),
		fx.Annotate(
			mapper.NewPlanMapper,
			fx.As(new(mapper.PlanMapperI)),
		),
		fx.Annotate(
			repository.NewSubscriptionRepository,
			fx.As(new(domain_repository.SubscriptionRepositoryI)),
		),
		fx.Annotate(
			repository.NewPlanRepository,
			fx.As(new(domain_repository.PlanRepositoryI)),
		),
	),
	fx.Invoke(
		router.SetupSubscriptionRoutes,
	),
)
