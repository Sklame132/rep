package games_postgres_repository

import (
	"time"

	"github.com/Sklame132/rep/internal/core/domain"
)

type GamesModel struct {
	ID        string
	FenStart  string
	FenEnd    string
	PlayerW   string
	PlayerB   string
	Type      string
	Mode      string
	Result    string
	History   any
	CreatedAt time.Time
}

func gameDomainsFromModels(games []GamesModel) []domain.Game {
	gameDomains := make([]domain.Game, len(games))

	for i, game := range games {
		gameDomains[i] = domain.NewGame(
			game.ID,
			game.FenStart,
			game.FenEnd,
			game.PlayerW,
			game.PlayerB,
			game.Type,
			game.Mode,
			game.Result,
			game.History,
			game.CreatedAt,
		)
	}

	return gameDomains
}
