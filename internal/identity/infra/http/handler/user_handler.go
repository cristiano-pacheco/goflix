package handler

import (
	"errors"
	"net/http"

	"github.com/cristiano-pacheco/goflix/internal/identity/application/usecase"
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/errs"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/http/dto"
	shared_errs "github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
	"github.com/cristiano-pacheco/goflix/internal/shared/sdk/http/request"
	"github.com/cristiano-pacheco/goflix/internal/shared/sdk/http/response"
)

type UserHandler struct {
	errorMapper         shared_errs.ErrorMapper
	userCreateUseCase   *usecase.UserCreateUseCase
	userUpdateUseCase   *usecase.UserUpdateUseCase
	userFindUseCase     *usecase.UserFindUseCase
	userActivateUseCase *usecase.UserActivateUseCase
}

func NewUserHandler(
	errorMapper shared_errs.ErrorMapper,
	userCreateUseCase *usecase.UserCreateUseCase,
	userUpdateUseCase *usecase.UserUpdateUseCase,
	userFindUseCase *usecase.UserFindUseCase,
	userActivateUseCase *usecase.UserActivateUseCase,
) *UserHandler {
	return &UserHandler{
		errorMapper,
		userCreateUseCase,
		userUpdateUseCase,
		userFindUseCase,
		userActivateUseCase,
	}
}

// @Summary		Create user
// @Description	Creates a new user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		request	body	dto.CreateUserRequest	true	"User data"
// @Success		201	{object}	response.Envelope[dto.CreateUserResponse]	"Successfully created user"
// @Failure		422	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "UserHandler.Create")
	defer span.End()

	var createUserRequest dto.CreateUserRequest
	if err := request.ReadJSON(w, r, &createUserRequest); err != nil {
		response.Error(w, err)
		return
	}

	input := usecase.UserCreateInput{
		Name:     createUserRequest.Name,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	output, err := h.userCreateUseCase.Execute(ctx, input)
	if err != nil {
		if errors.Is(err, errs.ErrEmailAlreadyInUse) {
			rError := h.errorMapper.MapCustomError(http.StatusBadRequest, err.Error())
			response.Error(w, rError)
			return
		}
		rError := h.errorMapper.Map(err)
		response.Error(w, rError)
		return
	}

	resData := dto.CreateUserResponse{
		UserID: output.UserID,
		Name:   output.Name,
		Email:  output.Email,
	}

	envelope := response.NewEnvelope(resData)
	response.JSON(w, http.StatusCreated, envelope, nil)
}

// @Summary		Update user
// @Description	Updates an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security 	BearerAuth
// @Param		id		path	integer		true	"User ID"
// @Param		request	body	dto.UpdateUserRequest	true	"User data"
// @Failure		422	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/me [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "UserHandler.Update")
	defer span.End()

	var updateUserRequest dto.UpdateUserRequest
	if err := request.ReadJSON(w, r, &updateUserRequest); err != nil {
		response.Error(w, err)
		return
	}

	userID, ok := r.Context().Value("user_id").(uint64)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	input := usecase.UserUpdateInput{
		UserID: userID,
		Name:   updateUserRequest.Name,
	}

	err := h.userUpdateUseCase.Execute(ctx, input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		response.Error(w, rError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary		Find user
// @Description	Finds an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security 	BearerAuth
// @Param		id		path	integer		true	"User ID"
// @Success		200	{object}	response.Envelope[dto.FindUserResponse]	"Successfully found user"
// @Failure		401	{object}	errs.Error	"Invalid credentials"
// @Failure		404	{object}	errs.Error	"User not found"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/me [get]
func (h *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "UserHandler.FindByID")
	defer span.End()

	userID, ok := r.Context().Value("user_id").(uint64)
	if !ok {
		response.JSON(w, http.StatusUnauthorized, nil, nil)
		return
	}

	input := usecase.UserFindInput{UserID: userID}
	output, err := h.userFindUseCase.Execute(ctx, input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		response.Error(w, rError)
		return
	}

	resData := dto.FindUserResponse{
		Name:  output.Name,
		Email: output.Email,
	}

	envelope := response.NewEnvelope(resData)
	response.JSON(w, http.StatusOK, envelope, nil)
}

// @Summary		Activate user
// @Description	Activates an existing user
// @Tags		Users
// @Accept		json
// @Produce		json
// @Param		request	body	dto.ActivateUserRequest	true	"User data"
// @Success		204		"Successfully activated user"
// @Failure		400	{object}	errs.Error	"Invalid request format or validation error"
// @Failure		500	{object}	errs.Error	"Internal server error"
// @Router		/api/v1/users/activate [post]
func (h *UserHandler) Activate(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Trace().StartSpan(r.Context(), "UserHandler.Activate")
	defer span.End()

	var req dto.ActivateUserRequest
	if err := request.ReadJSON(w, r, &req); err != nil {
		response.Error(w, err)
		return
	}

	input := usecase.UserActivateInput{Token: req.Token}
	err := h.userActivateUseCase.Execute(ctx, input)
	if err != nil {
		rError := h.errorMapper.Map(err)
		response.Error(w, rError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
