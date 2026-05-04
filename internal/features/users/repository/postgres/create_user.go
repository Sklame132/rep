package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	INSERT INTO rep.users (username, password, first_name, last_name, address, email, phone_number, rating, role, image_url) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING *;
	`

	row := r.pool.QueryRow(ctx, query, user.Username, user.Password, user.FirstName, user.LastName, user.Address, user.Email, user.PhoneNumber, user.Rating, user.Role, user.ImageURL)

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

	return userDomain, nil
}
