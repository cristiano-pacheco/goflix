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
		NewFacade,
		usecase.NewSubscriptionCreateUseCase,
		handler.NewSubscriptionHandler,
		domain_mapper.NewEndDateMapper,
		mapper.NewSubscriptionMapper,
		mapper.NewPlanMapper,
		fx.Annotate(
			repository.NewSubscriptionRepository,
			fx.As(new(domain_repository.SubscriptionRepository)),
		),
		fx.Annotate(
			repository.NewPlanRepository,
			fx.As(new(domain_repository.PlanRepository)),
		),
	),
	fx.Invoke(
		router.SetupSubscriptionRoutes,
	),
)
