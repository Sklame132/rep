package users_transport_http

import (
	"net/http"

	core_logger "github.com/Sklame132/rep/internal/core/logger"
	core_http_request "github.com/Sklame132/rep/internal/core/transport/http/request"
	core_http_response "github.com/Sklame132/rep/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userLogin, err := core_http_request.GetStringPathValue(r, "login")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userLogin path value",
		)

		return
	}

	if err := h.usersService.DeleteUser(ctx, userLogin); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete user",
		)

		return
	}

	responseHandler.NoContentResponse()
}
