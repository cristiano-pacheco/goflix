package handler

import (
	"encoding/json"
	"net/http"

	"github.com/cristiano-pacheco/goflix/internal/identity/application/usecase"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/dto"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
	"github.com/cristiano-pacheco/goflix/internal/shared/sdk/http/response"
)

type AuthHandler struct {
	errorMapper          errs.ErrorMapper
	tokenGenerateUseCase usecase.TokenGenerateUseCase
}

func NewAuthHandler(
	errorMapper errs.ErrorMapper,
	tokenGenerateUseCase usecase.TokenGenerateUseCase,
) *AuthHandler {
	return &AuthHandler{errorMapper, tokenGenerateUseCase}
}

// @Summary		Generate authentication token
// @Description	Authenticates user credentials and returns an access token
// @Tags		Authentication
// @Accept		json
// @Produce		json
// @Param		request	body	dto.GenerateTokenRequest	true	"Login credentials (email and password)"
// @Success		200	{object}	response.Envelope[dto.GenerateTokenResponse]	"Successfully generated token"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/auth/token [post]
func (h *AuthHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "AuthHandler.GenerateToken")
	defer span.End()

	var request dto.GenerateTokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Error(w, err)
		return
	}

	input := usecase.TokenGenerateInput{
		Email:    request.Email,
		Password: request.Password,
	}

	output, err := h.tokenGenerateUseCase.Execute(ctx, input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		response.Error(w, rError)
		return
	}

	generateTokenResponse := dto.GenerateTokenResponse{Token: output.Token}
	envelope := response.NewEnvelope(generateTokenResponse)
	response.JSON(w, http.StatusOK, envelope, nil)
}
