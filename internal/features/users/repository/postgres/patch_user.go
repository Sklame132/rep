package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	login string,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	UPDATE rep.users
	SET
		username=$1,
		password=$2,
		first_name=$3,
		last_name=$4,
		address=$5,
		email=$6,
		phone_number=$7,
		rating=$8,
		role=$9
	WHERE id=uuid_or_null($10) OR username=$10
	RETURNING *;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.Username,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Address,
		user.Email,
		user.PhoneNumber,
		user.Rating,
		user.Role,
		login,
	)

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
