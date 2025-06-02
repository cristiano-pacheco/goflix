package handler

import (
	"errors"
	"net/http"

	"github.com/cristiano-pacheco/goflix/internal/billing/application/usecase"
	"github.com/cristiano-pacheco/goflix/internal/billing/domain/errs"
	"github.com/cristiano-pacheco/goflix/internal/billing/infra/http/dto"
	shared_errs "github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
	"github.com/cristiano-pacheco/goflix/internal/shared/sdk/http/request"
	"github.com/cristiano-pacheco/goflix/internal/shared/sdk/http/response"
)

type SubscriptionHandler struct {
	errorMapper               shared_errs.ErrorMapper
	createSubscriptionUseCase *usecase.CreateSubscriptionUseCase
}

func NewSubscriptionHandler(
	errorMapper shared_errs.ErrorMapper,
	createSubscriptionUseCase *usecase.CreateSubscriptionUseCase,
) *SubscriptionHandler {
	return &SubscriptionHandler{
		errorMapper,
		createSubscriptionUseCase,
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
