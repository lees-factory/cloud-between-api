package user

import (
	"context"
)

type UserUpdater struct {
	userRepository UserRepository
}

func NewUserUpdater(userRepository UserRepository) *UserUpdater {
	return &UserUpdater{userRepository: userRepository}
}

func (u *UserUpdater) UpdateLastLogin(ctx context.Context, email string) error {
	return u.userRepository.UpdateLastLogin(ctx, email)
}
