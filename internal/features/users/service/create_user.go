package users_service

import (
	"context"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(password)
	
	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
