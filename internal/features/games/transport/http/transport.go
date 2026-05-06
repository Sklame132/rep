package games_transport_http

import (
	"context"
	"net/http"

	"github.com/Sklame132/rep/internal/core/domain"
	core_http_server "github.com/Sklame132/rep/internal/core/transport/http/server"
)

type GamesHTTPHandler struct {
	gamesService GamesService
}

type GamesService interface {
	CreateGame(
		ctx context.Context,
		game domain.Game,
	) (domain.Game, error)

	GetGames(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.Game, error)
}

func NewGamesHTTPHandler(
	gamesService GamesService,
) *GamesHTTPHandler {
	return &GamesHTTPHandler{
		gamesService: gamesService,
	}
}

func (h *GamesHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/games",
			Handler: h.CreateGame,
		},
		{
			Method: http.MethodGet,
			Path: "/games",
			Handler: h.GetGames,
		},
	}
}
