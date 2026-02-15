package repository

import (
	"context"

	"gorm.io/gorm"
	domainuser "io.lees.cloud-between/core/core-domain/user"
	"io.lees.cloud-between/storage/db-core/entity"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserCoreRepository(db *gorm.DB) domainuser.UserRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Save(ctx context.Context, u *domainuser.User) error {
	var socialProvider *string
	if u.SocialProvider != nil {
		p := string(*u.SocialProvider)
		socialProvider = &p
	}

	e := &entity.UserEntity{
		BaseEntity: entity.BaseEntity{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
		},
		SocialID:        u.SocialID,
		SocialProvider:  socialProvider,
		Email:           u.Email,
		PasswordHash:    u.PasswordHash,
		ProfileImageURL: u.ProfileImageURL,
		IsPaid:          u.IsPaid,
		LastLoginAt:     u.LastLoginAt,
	}
	return r.db.WithContext(ctx).Save(e).Error
}

func (r *UserGormRepository) FindByEmail(ctx context.Context, email string) (*domainuser.User, error) {
	var e entity.UserEntity
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&e).Error
	if err != nil {
		return nil, err
	}
	return toUserDomain(e), nil
}

func (r *UserGormRepository) FindBySocialIDAndProvider(ctx context.Context, socialID string, provider domainuser.SocialProvider) (*domainuser.User, error) {
	var e entity.UserEntity
	err := r.db.WithContext(ctx).
		Where("social_id = ? AND social_provider = ?", socialID, string(provider)).
		First(&e).Error
	if err != nil {
		return nil, err
	}
	return toUserDomain(e), nil
}

func (r *UserGormRepository) UpdateLastLogin(ctx context.Context, email string) error {
	return r.db.WithContext(ctx).Model(&entity.UserEntity{}).
		Where("email = ?", email).
		Update("last_login_at", gorm.Expr("NOW()")).Error
}

func toUserDomain(e entity.UserEntity) *domainuser.User {
	var provider *domainuser.SocialProvider
	if e.SocialProvider != nil {
		p := domainuser.SocialProvider(*e.SocialProvider)
		provider = &p
	}

	return &domainuser.User{
		ID:              e.ID,
		SocialID:        e.SocialID,
		SocialProvider:  provider,
		Email:           e.Email,
		PasswordHash:    e.PasswordHash,
		ProfileImageURL: e.ProfileImageURL,
		IsPaid:          e.IsPaid,
		LastLoginAt:     e.LastLoginAt,
		CreatedAt:       e.CreatedAt,
	}
}
