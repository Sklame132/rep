package games_transport_http

import (
	"net/http"

	core_logger "github.com/Sklame132/rep/internal/core/logger"
	core_http_request "github.com/Sklame132/rep/internal/core/transport/http/request"
	core_http_response "github.com/Sklame132/rep/internal/core/transport/http/response"
)

type GetGamesResponse []GameDTOResponse

func (h *GamesHTTPHandler) GetGames(w http.ResponseWriter, r *http.Request) {
	const (
		usernameQueryParamKey = "username"
	)

	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	limit, offset, err := core_http_request.GetLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get `limit` / `offset` query param")

		return
	}

	usernameParam := r.URL.Query().Get(usernameQueryParamKey)
	var username *string

	if usernameParam != "" {
		username = &usernameParam
	}

	gameDomains, err := h.gamesService.GetGames(ctx, limit, offset, username)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get games")

		return
	}

	response := GetGamesResponse(gamesDTOFromDomains(gameDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}