package business

import (
	"context"
	"errors"

	"io.lees.cloud-between/core/core-domain/user/concept"
	"io.lees.cloud-between/core/core-domain/user/implement"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userAppender *implement.UserAppender
	userFinder   *implement.UserFinder
	userUpdater  *implement.UserUpdater
}

func NewUserService(
	userAppender *implement.UserAppender,
	userFinder *implement.UserFinder,
	userUpdater *implement.UserUpdater,
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
	user := concept.NewUser(email, &pwHash)

	return s.userAppender.Append(ctx, user)
}

func (s *UserService) Login(ctx context.Context, email, password string) (*concept.User, error) {
	user, err := s.userFinder.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user.PasswordHash == nil {
		return nil, errors.New("user has no password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	_ = s.userUpdater.UpdateLastLogin(ctx, email)

	return user, nil
}
