package users_transport_http

import (
	"net/http"

	core_logger "github.com/Sklame132/rep/internal/core/logger"
	core_http_response "github.com/Sklame132/rep/internal/core/transport/http/response"
	core_http_utils "github.com/Sklame132/rep/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userLogin, err := core_http_utils.GetStringPathValue(r, "login")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userLogin path value",
		)

		return
	}

	user, err := h.usersService.GetUser(ctx, userLogin)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)

		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}
