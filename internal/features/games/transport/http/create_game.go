package games_transport_http

import (
	"net/http"

	"github.com/Sklame132/rep/internal/core/domain"
	core_logger "github.com/Sklame132/rep/internal/core/logger"
	core_http_request "github.com/Sklame132/rep/internal/core/transport/http/request"
	core_http_response "github.com/Sklame132/rep/internal/core/transport/http/response"
)

type CreateGameRequest struct {
	FenStart string `json:"fen_start" validate:"required,max=80"`
	FenEnd string `json:"fen_end" validate:"required,max=80"`
	PlayerW string `json:"player_w" validate:"required,max=16"`
	PlayerB string `json:"player_b" validate:"required,max=16"`
	Type string `json:"type" validate:"required,max=7"`
	Mode string `json:"mode" validate:"required,max=16"`
	Result string `json:"result" validate:"required,max=5"`
	History *string `json:"history"`
}

type CreateGameResponse GameDTOResponse

func (h *GamesHTTPHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoke CreateGame handler")

	var request CreateGameRequest
	if err := core_http_request.DecodeAndValidationRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		
		return
	}

	gameDomain := domainFromDTO(request)

	gameDomain, err := h.gamesService.CreateGame(ctx,
	gameDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create game")

		return
	}

	response := CreateGameResponse(gameDTOFromDomain(gameDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateGameRequest) domain.Game {
	return domain.NewGameUnitialized(
		dto.FenStart,
		dto.FenEnd,
		dto.PlayerW,
		dto.PlayerB,
		dto.Type,
		dto.Mode,
		dto.Result,
		dto.History,
	)
}