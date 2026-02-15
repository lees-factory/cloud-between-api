package implement

import (
	"context"
	"time"

	"io.lees.cloud-between/storage/db-core/user"
)

type UserUpdater struct {
	userRepository user.UserRepository
}

func NewUserUpdater(userRepository user.UserRepository) *UserUpdater {
	return &UserUpdater{userRepository: userRepository}
}

func (u *UserUpdater) UpdateLastLogin(ctx context.Context, email string) error {
	entity, err := u.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	entity.LastLoginAt = time.Now()
	return u.userRepository.Save(ctx, entity)
}
