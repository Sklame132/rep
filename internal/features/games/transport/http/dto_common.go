package games_transport_http

import (
	"time"

	"github.com/Sklame132/rep/internal/core/domain"
)

type GameDTOResponse struct {
	ID        string    `json:"id"`
	FenStart  string    `json:"fen_start"`
	FenEnd    string    `json:"fen_end"`
	PlayerW   string    `json:"player_w"`
	PlayerB   string    `json:"player_b"`
	Type      string    `json:"type"`
	Mode      string    `json:"mode"`
	Result    string    `json:"result"`
	History   any       `json:"history"`
	CreatedAt time.Time `json:"created_at"`
}

func gameDTOFromDomain(game domain.Game) GameDTOResponse {
	return GameDTOResponse{
		ID: game.ID,
		FenStart: game.FenStart,
		FenEnd: game.FenEnd,
		PlayerW: game.PlayerW,
		PlayerB: game.PlayerB,
		Type: game.Type,
		Mode: game.Mode,
		Result: game.Result,
		History: game.History,
		CreatedAt: game.CreatedAt,
	}
}

func gamesDTOFromDomains(games []domain.Game) []GameDTOResponse {
	gamesDTO := make([]GameDTOResponse, len(games))

	for i, game := range games {
		gamesDTO[i] = gameDTOFromDomain(game)
	}

	return gamesDTO
}