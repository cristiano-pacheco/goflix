package router

import (
	"net/http"

	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/handler"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/http/router"
)

func SetupAuthRoutes(r *router.Router, authHandler *handler.AuthHandler) {
	router := r.Router()
	router.HandlerFunc(http.MethodPost, "/api/v1/auth/token", authHandler.GenerateToken)
}
