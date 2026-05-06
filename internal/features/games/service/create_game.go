package games_service

import (
	"context"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
)

func (s *GamesService) CreateGame(
	ctx context.Context,
	game domain.Game,
) (domain.Game, error) {
	if err := game.Validate(); err != nil {
		return domain.Game{}, fmt.Errorf("validate game domain: %w", err)
	}

	game, err := s.gamesRepository.CreateGame(ctx, game)
	if err != nil {
		return domain.Game{}, fmt.Errorf("create game: %w", err)
	}

	return game, nil
}
