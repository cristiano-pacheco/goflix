package router

import (
	"net/http"

	"github.com/cristiano-pacheco/goflix/internal/billing/infra/http/handler"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/http/middleware"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/http/router"
)

func SetupSubscriptionRoutes(
	r *router.Router,
	subscriptionHandler *handler.SubscriptionHandler,
	authMiddleware *middleware.AuthMiddleware,
) {
	router := r.Router()
	router.HandlerFunc(
		http.MethodPost,
		"/api/v1/subscriptions",
		authMiddleware.Middleware(subscriptionHandler.Create),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/api/v1/subscriptions",
		authMiddleware.Middleware(subscriptionHandler.FindByUserID),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/api/v1/subscriptions/active",
		authMiddleware.Middleware(subscriptionHandler.IsUserSubscriptionActive),
	)
}
