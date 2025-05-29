package router

import (
	"net/http"

	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/handler"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/middleware"
)

func SetupUserRoutes(
	r *Router,
	userHandler *handler.UserHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	router := r.Router()
	router.HandlerFunc(http.MethodPost, "/api/v1/users", userHandler.Create)
	router.HandlerFunc(http.MethodPost, "/api/v1/users/activate", userHandler.Activate)
	router.HandlerFunc(http.MethodGet, "/api/v1/users/me", authMiddleware.Middleware(userHandler.FindByID))
	router.HandlerFunc(http.MethodPut, "/api/v1/users/me", authMiddleware.Middleware(userHandler.Update))
}
