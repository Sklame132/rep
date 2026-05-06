package games_service

import (
	"context"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
	core_errors "github.com/Sklame132/rep/internal/core/errors"
)

func (s *GamesService) GetGames(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Game, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	games, err := s.gamesRepository.GetGames(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get games from repository: %w", err)
	}

	return games, nil
}
