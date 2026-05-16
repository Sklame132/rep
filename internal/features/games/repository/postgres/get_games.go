package games_postgres_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Sklame132/rep/internal/core/domain"
)

func (r *GamesRepository) GetGames(
	ctx context.Context,
	limit *int,
	offset *int,
	username *string,
) ([]domain.Game, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
		SELECT * FROM rep.games
	`)

	args := []any{limit, offset}
	conditions := []string{}

	if username != nil {
		conditions = append(conditions, fmt.Sprintf("player_w=$%d OR player_b=$%d", len(args)+1, len(args)+1))
		args = append(args, username)
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString("WHERE " + strings.Join(conditions, ""))
	}
	queryBuilder.WriteString(`
		ORDER BY created_at DESC
		LIMIT $1
		OFFSET $2;
	`)

	rows, err := r.pool.Query(
		ctx,
		queryBuilder.String(),
		args...
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
