package users_postgres_repository

import (
	"time"

	"github.com/Sklame132/rep/internal/core/domain"
)

type UserModel struct {
	ID          string
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Address     *string
	Email       *string
	PhoneNumber *string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	Rating      *int16
	Role 		*string
	ImageURL    *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Username,
			user.Password,
			user.FirstName,
			user.LastName,
			user.Address,
			user.Email,
			user.PhoneNumber,
			user.CreatedAt,
			user.UpdatedAt,
			user.Rating,
			user.Role,
			user.ImageURL,
		)
	}

	return userDomains
}