package users_transport_http

import (
	"net/http"

	"github.com/Sklame132/rep/internal/core/domain"
	core_logger "github.com/Sklame132/rep/internal/core/logger"
	core_http_request "github.com/Sklame132/rep/internal/core/transport/http/request"
	core_http_response "github.com/Sklame132/rep/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	Username    string  `json:"username" validate:"required,min=3,max=16"`
	Password    string  `json:"password" validate:"required,min=8,max=64"`
	FirstName   string  `json:"first_name" validate:"required,max=32"`
	LastName    string  `json:"last_name" validate:"required,max=32"`
	Address     *string `json:"address"`
	Email       *string `json:"email" validate:"omitempty,max=32"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=16,startswith=+"`
	Rating      *int16  `json:"rating"`
	Role        *string `json:"role" validate:"omitempty,max=16"`
}

type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke CreateUser handler")

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidationRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")

		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(
		dto.Username,
		dto.Password,
		dto.FirstName,
		dto.LastName,
		dto.Address,
		dto.Email,
		dto.PhoneNumber,
		dto.Rating,
		dto.Role,
	)
}
