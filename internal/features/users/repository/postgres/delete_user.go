package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Sklame132/rep/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(
	ctx context.Context,
	login string,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
		DELETE FROM rep.users
		WHERE id=uuid_or_null($1) OR username=$1
	`

	cmdTag, err := r.pool.Exec(ctx, query, login)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user with id or username – %s: %w", login, core_errors.ErrNotFound)
	}

	return nil
}
