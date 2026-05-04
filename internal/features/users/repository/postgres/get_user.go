package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
	core_errors "github.com/Sklame132/rep/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	login string,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
		SELECT * FROM rep.users
		WHERE id=uuid_or_null($1) OR username=$1;
	`

	row := r.pool.QueryRow(ctx, query, login)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Username,
		&userModel.Password,
		&userModel.FirstName,
		&userModel.LastName,
		&userModel.Address,
		&userModel.Email,
		&userModel.PhoneNumber,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
		&userModel.Rating,
		&userModel.Role,
		&userModel.ImageURL,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id or username – %s: %w", login, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
			userModel.ID,
			userModel.Username,
			userModel.Password,
			userModel.FirstName,
			userModel.LastName,
			userModel.Address,
			userModel.Email,
			userModel.PhoneNumber,
			userModel.CreatedAt,
			userModel.UpdatedAt,
			userModel.Rating,
			userModel.Role,
			userModel.ImageURL,
	)

	return userDomain, err
}
