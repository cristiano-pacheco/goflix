package billing

import (
	"go.uber.org/fx"

	"github.com/cristiano-pacheco/goflix/internal/billing/application/usecase"
	domain_mapper "github.com/cristiano-pacheco/goflix/internal/billing/domain/mapper"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/http/handler"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/http/router"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/persistence/gorm/repository"
	"github.com/cristiano-pacheco/goflix/internal/billing/ports"
)

var Module = fx.Module(
	"billing",
	fx.Provide(
		usecase.NewSubscriptionCreateUseCase,
		handler.NewSubscriptionHandler,
		fx.Annotate(
			NewFacade,
			fx.As(new(FacadeI)),
		),
		fx.Annotate(
			domain_mapper.NewEndDateMapper,
			fx.As(new(domain_mapper.EndDateMapperI)),
		),
		mapper.NewSubscriptionMapper,
		mapper.NewPlanMapper,
		fx.Annotate(
			repository.NewSubscriptionRepository,
			fx.As(new(ports.SubscriptionRepositoryI)),
		),
		fx.Annotate(
			repository.NewPlanRepository,
			fx.As(new(ports.PlanRepositoryI)),
		),
	),
	fx.Invoke(
		router.SetupSubscriptionRoutes,
	),
)
