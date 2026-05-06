package games_service

import (
	"context"

	"github.com/Sklame132/rep/internal/core/domain"
)

type GamesService struct {
	gamesRepository GamesRepository
}

type GamesRepository interface {
	CreateGame(
		ctx context.Context,
		game domain.Game,
	) (domain.Game, error)

	GetGames(
		ctx context.Context,
		limit *int,
		offset *int,
		username *string,
	) ([]domain.Game, error)
}

func NewGamesService(
	gamesRepository GamesRepository,
) *GamesService {
	return &GamesService {
		gamesRepository: gamesRepository,
	}
}