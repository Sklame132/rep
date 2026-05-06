package games_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
	core_errors "github.com/Sklame132/rep/internal/core/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *GamesRepository) CreateGame(
	ctx context.Context,
	game domain.Game,
) (domain.Game, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
		INSERT INTO rep.games (fen_start, fen_end, player_w, player_b, type, mode, result, history)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING *;
	`

	row := r.pool.QueryRow(ctx, query, game.FenStart, game.FenEnd, game.PlayerW, game.PlayerB, game.Type, game.Mode, game.Result, game.History)

	var gameModel GamesModel

	err := row.Scan(
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

	const pgxViolatesForeignKeyErrorCode = "23503"

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgxViolatesForeignKeyErrorCode {
				return domain.Game{}, fmt.Errorf("violates foreign key: %v: user with logins: `%s` or `%s`: %w", err, game.PlayerW, game.PlayerB, core_errors.ErrNotFound)
			}
		}
		return domain.Game{}, fmt.Errorf("scan error: %w", err)
	}

	gameDomain := domain.NewGame(
		gameModel.ID,
		gameModel.FenStart,
		gameModel.FenEnd, 
		gameModel.PlayerW,
		gameModel.PlayerB,
		gameModel.Type,
		gameModel.Mode,
		gameModel.Result,
		gameModel.History,
		gameModel.CreatedAt,
	)

	return gameDomain, nil
}
