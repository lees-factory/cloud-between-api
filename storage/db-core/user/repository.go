package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(ctx context.Context, entity *UserEntity) error
	FindByEmail(ctx context.Context, email string) (*UserEntity, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, entity *UserEntity) error {
	return r.db.WithContext(ctx).Save(entity).Error
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*UserEntity, error) {
	var entity UserEntity
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
