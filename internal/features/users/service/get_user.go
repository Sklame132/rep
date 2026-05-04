package users_service

import (
	"context"
	"fmt"

	"github.com/Sklame132/rep/internal/core/domain"
)

func (s *UsersService) GetUser(
	ctx context.Context,
	login string,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, login)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}
