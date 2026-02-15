package implement

import (
	"context"

	"io.lees.cloud-between/core/core-domain/user/concept"
	"io.lees.cloud-between/storage/db-core/user"
)

type UserFinder struct {
	userRepository user.UserRepository
}

func NewUserFinder(userRepository user.UserRepository) *UserFinder {
	return &UserFinder{userRepository: userRepository}
}

func (f *UserFinder) FindByEmail(ctx context.Context, email string) (*concept.User, error) {
	entity, err := f.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &concept.User{
		ID:              entity.ID,
		GoogleID:        entity.GoogleID,
		Email:           entity.Email,
		PasswordHash:    entity.PasswordHash,
		ProfileImageURL: entity.ProfileImageURL,
		IsPaid:          entity.IsPaid,
		LastLoginAt:     entity.LastLoginAt,
		CreatedAt:       entity.CreatedAt,
	}, nil
}
