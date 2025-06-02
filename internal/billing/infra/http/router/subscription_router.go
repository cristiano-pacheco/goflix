package router

import (
	"net/http"

	"github.com/cristiano-pacheco/goflix/internal/billing/infra/http/handler"
)

func SetupSubscriptionRoutes(
	r *Router,
	subscriptionHandler *handler.SubscriptionHandler,
) {
	router := r.Router()
	router.HandlerFunc(http.MethodPost, "/api/v1/subscriptions", subscriptionHandler.Create)
}
