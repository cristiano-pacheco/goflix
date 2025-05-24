package modules

import (
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/config"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/database"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/httpserver"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/jwt"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/logger"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/mailer"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/registry"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/translator"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/validator"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"shared/modules",
	config.Module,
	database.Module,
	validator.Module,
	translator.Module,
	httpserver.Module,
	logger.Module,
	registry.Module,
	jwt.Module,
	mailer.Module,
	errs.Module,
)
