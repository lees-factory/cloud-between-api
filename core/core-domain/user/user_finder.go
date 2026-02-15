package user

import (
	"context"
)

type UserFinder struct {
	userRepository UserRepository
}

func NewUserFinder(userRepository UserRepository) *UserFinder {
	return &UserFinder{userRepository: userRepository}
}

func (f *UserFinder) FindByEmail(ctx context.Context, email string) (*User, error) {
	return f.userRepository.FindByEmail(ctx, email)
}

func (f *UserFinder) FindBySocialIDAndProvider(ctx context.Context, socialID string, provider SocialProvider) (*User, error) {
	return f.userRepository.FindBySocialIDAndProvider(ctx, socialID, provider)
}
