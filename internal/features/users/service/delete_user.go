package users_service

import (
	"context"
	"fmt"
)

func (s *UsersService) DeleteUser(
	ctx context.Context,
	login string,
) error {
	if err := s.usersRepository.DeleteUser(ctx, login); err != nil {
		return fmt.Errorf("delete user from repository: %w", err)
	}

	return nil
}
