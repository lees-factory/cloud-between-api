package user

import (
	"context"
)

type UserAppender struct {
	userRepository UserRepository
}

func NewUserAppender(userRepository UserRepository) *UserAppender {
	return &UserAppender{userRepository: userRepository}
}

func (a *UserAppender) Append(ctx context.Context, u *User) error {
	return a.userRepository.Save(ctx, u)
}
