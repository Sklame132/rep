package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
)

func (r *UsersRepository) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
		SELECT * FROM rep.users
		ORDER BY id ASC
		LIMIT $1
		OFFSET $2
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select users: %w", err)
	}
	defer rows.Close()

	var userModels []UserModel

	for rows.Next() {
		var userModel UserModel

		err := rows.Scan(
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
			return nil, fmt.Errorf("scan error: %w", err)
		}

		userModels = append(userModels, userModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	userDomains := userDomainsFromModels(userModels)
	
	return userDomains, nil
}
