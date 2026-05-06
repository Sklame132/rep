package games_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
)

func (r *GamesRepository) GetGames(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.Game, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
		SELECT * FROM rep.games
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2;
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select games: %w", err)
	}
	defer rows.Close()

	var gameModels []GamesModel

	for rows.Next() {
		var gameModel GamesModel

		err := rows.Scan(
			&gameModel.ID,
			&gameModel.FenStart,
			&gameModel.FenEnd,
			&gameModel.PlayerW,
			&gameModel.PlayerB,
			&gameModel.Type,
			&gameModel.Mode,
			&gameModel.Result,
			&gameModel.History,
			&gameModel.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		gameModels = append(gameModels, gameModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	gameDomains := gameDomainsFromModels(gameModels)

	return gameDomains, nil
}
