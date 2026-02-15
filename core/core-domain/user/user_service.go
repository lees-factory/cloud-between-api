package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userAppender *UserAppender
	userFinder   *UserFinder
	userUpdater  *UserUpdater
}

func NewUserService(
	userAppender *UserAppender,
	userFinder *UserFinder,
	userUpdater *UserUpdater,
) *UserService {
	return &UserService{
		userAppender: userAppender,
		userFinder:   userFinder,
		userUpdater:  userUpdater,
	}
}

func (s *UserService) Signup(ctx context.Context, email, password string) error {
	existing, _ := s.userFinder.FindByEmail(ctx, email)
	if existing != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	pwHash := string(hashedPassword)
	u := NewUser(email, &pwHash)

	return s.userAppender.Append(ctx, u)
}

func (s *UserService) Login(ctx context.Context, email, password string) (*User, error) {
	u, err := s.userFinder.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if u.PasswordHash == nil {
		return nil, errors.New("user has no password (try social login)")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*u.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	_ = s.userUpdater.UpdateLastLogin(ctx, email)

	return u, nil
}

func (s *UserService) LoginBySocial(ctx context.Context, socialID string, provider SocialProvider, email string) (u *User, isNew bool, err error) {
	if !provider.IsValid() {
		return nil, false, errors.New("unsupported social provider")
	}

	// 1. social_id + provider로 기존 유저 조회
	u, err = s.userFinder.FindBySocialIDAndProvider(ctx, socialID, provider)
	if err == nil {
		_ = s.userUpdater.UpdateLastLogin(ctx, u.Email)
		return u, false, nil
	}

	// 2. 없으면 자동 회원가입
	u = NewSocialUser(email, socialID, provider)
	if err := s.userAppender.Append(ctx, u); err != nil {
		return nil, false, err
	}

	return u, true, nil
}
