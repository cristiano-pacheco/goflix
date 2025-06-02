package handler

import (
	"errors"
	"net/http"

	"github.com/cristiano-pacheco/goflix/internal/billing/application/usecase"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/repository"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/http/dto"
	shared_errs "github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
	"github.com/cristiano-pacheco/goflix/internal/shared/sdk/http/request"
	"github.com/cristiano-pacheco/goflix/internal/shared/sdk/http/response"
)

type SubscriptionHandler struct {
	errorMapper               shared_errs.ErrorMapper
	createSubscriptionUseCase *usecase.CreateSubscriptionUseCase
	subscriptionRepository    repository.SubscriptionRepository
}

func NewSubscriptionHandler(
	errorMapper shared_errs.ErrorMapper,
	createSubscriptionUseCase *usecase.CreateSubscriptionUseCase,
	subscriptionRepository repository.SubscriptionRepository,
) *SubscriptionHandler {
	return &SubscriptionHandler{
		errorMapper,
		createSubscriptionUseCase,
		subscriptionRepository,
	}
}

// @Summary		Create subscription
// @Description	Creates a new subscription for the authenticated user
// @Tags		Subscriptions
// @Accept		json
// @Produce		json
// @Security 	BearerAuth
// @Param		request	body	dto.CreateSubscriptionRequest	true	"Subscription data"
// @Success		201	{object}	response.Envelope[dto.CreateSubscriptionResponse]	"Successfully created subscription"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		422	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/subscriptions [post]
func (h *SubscriptionHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "SubscriptionHandler.CreateSubscription")
	defer span.End()

	var createSubscriptionRequest dto.CreateSubscriptionRequest
	if err := request.ReadJSON(w, r, &createSubscriptionRequest); err != nil {
		response.Error(w, err)
		return
	}

	userID := request.GetUserID(r)
	if userID == 0 {
		response.JSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	input := usecase.CreateSubscriptionInput{
		PlanID: createSubscriptionRequest.PlanID,
		UserID: userID,
	}

	output, err := h.createSubscriptionUseCase.Execute(ctx, input)
	if err != nil {
		if errors.Is(err, errs.ErrPlanNotFound) {
			rError := h.errorMapper.MapCustomError(http.StatusBadRequest, err.Error())
			response.Error(w, rError)
			return
		}
		if errors.Is(err, errs.ErrUserAlreadyHasActiveSubscription) {
			rError := h.errorMapper.MapCustomError(http.StatusBadRequest, err.Error())
			response.Error(w, rError)
			return
		}
		rError := h.errorMapper.Map(err)
		response.Error(w, rError)
		return
	}

	resData := dto.CreateSubscriptionResponse{
		SubscriptionID: output.SubscriptionID,
		UserID:         output.UserID,
		PlanID:         output.PlanID,
		Status:         output.Status,
		StartDate:      output.StartDate,
		EndDate:        output.EndDate,
		AutoRenew:      output.AutoRenew,
	}

	envelope := response.NewEnvelope(resData)
	response.JSON(w, http.StatusCreated, envelope, nil)
}

// @Summary		List user subscriptions
// @Description	Retrieves all subscriptions for the authenticated user
// @Tags		Subscriptions
// @Accept		json
// @Produce		json
// @Security 	BearerAuth
// @Success		200	{object}	response.Envelope[dto.ListSubscriptionsResponse]	"Successfully retrieved subscriptions"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/subscriptions [get]
func (h *SubscriptionHandler) FindByUserID(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "SubscriptionHandler.FindByUserID")
	defer span.End()

	userID := request.GetUserID(r)
	if userID == 0 {
		response.JSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	subscriptions, err := h.subscriptionRepository.FindByUserID(ctx, userID)
	if err != nil {
		rError := h.errorMapper.Map(err)
		response.Error(w, rError)
		return
	}

	subscriptionResponses := make([]dto.SubscriptionResponse, 0, len(subscriptions))
	for _, subscription := range subscriptions {
		status := subscription.Status()
		subscriptionResponse := dto.SubscriptionResponse{
			SubscriptionID: subscription.ID(),
			UserID:         subscription.UserID(),
			PlanID:         subscription.PlanID(),
			Status:         status.String(),
			StartDate:      subscription.StartDate(),
			EndDate:        subscription.EndDate(),
			AutoRenew:      subscription.AutoRenew(),
		}
		subscriptionResponses = append(subscriptionResponses, subscriptionResponse)
	}

	resData := dto.ListSubscriptionsResponse{
		Subscriptions: subscriptionResponses,
	}

	envelope := response.NewEnvelope(resData)
	response.JSON(w, http.StatusOK, envelope, nil)
}

// @Summary		Check user subscription status
// @Description	Checks if the authenticated user has an active subscription
// @Tags		Subscriptions
// @Accept		json
// @Produce		json
// @Security 	BearerAuth
// @Success		200	{object}	response.Envelope[dto.IsUserSubscriptionActiveResponse]	"Successfully retrieved subscription status"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/subscriptions/active [get]
func (h *SubscriptionHandler) IsUserSubscriptionActive(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "SubscriptionHandler.IsUserSubscriptionActive")
	defer span.End()

	userID := request.GetUserID(r)
	if userID == 0 {
		response.JSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	_, err := h.subscriptionRepository.FindActiveSubscriptionByUserID(ctx, userID)
	isActive := err == nil

	resData := dto.IsUserSubscriptionActiveResponse{
		IsActive: isActive,
	}

	envelope := response.NewEnvelope(resData)
	response.JSON(w, http.StatusOK, envelope, nil)
}
