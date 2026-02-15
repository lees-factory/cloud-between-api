package implement

import (
	"context"

	"io.lees.cloud-between/core/core-domain/user/concept"
	"io.lees.cloud-between/storage/db-core/user"
)

type UserAppender struct {
	userRepository user.UserRepository
}

func NewUserAppender(userRepository user.UserRepository) *UserAppender {
	return &UserAppender{userRepository: userRepository}
}

func (a *UserAppender) Append(ctx context.Context, u *concept.User) error {
	entity := &user.UserEntity{
		ID:              u.ID,
		GoogleID:        u.GoogleID,
		Email:           u.Email,
		PasswordHash:    u.PasswordHash,
		ProfileImageURL: u.ProfileImageURL,
		IsPaid:          u.IsPaid,
		LastLoginAt:     u.LastLoginAt,
		CreatedAt:       u.CreatedAt,
	}
	return a.userRepository.Save(ctx, entity)
}
