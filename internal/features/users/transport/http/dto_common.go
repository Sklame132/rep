package users_transport_http

import (
	"time"

	"github.com/Sklame132/rep/internal/core/domain"
)

type UserDTOResponse struct {
	ID          string     `json:"id"`
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Address     *string    `json:"address"`
	Email       *string    `json:"email"`
	PhoneNumber *string    `json:"phone_number"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Rating      *int16      `json:"rating"`
	Role        *string    `json:"role"`
	ImageURL    *string    `json:"image_url"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Address:     user.Address,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Rating:      user.Rating,
		Role:        user.Role,
		ImageURL:    user.ImageURL,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
