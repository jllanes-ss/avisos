package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/jllanes-ss/avisos/account/domain"
	"github.com/jllanes-ss/avisos/account/domain/apperrors"
)

type userUseCase struct {
	UserRepository domain.UserRepository
}

type UUCConfig struct {
	UserRepository domain.UserRepository
}

// NewUserService will create new an userService object representation of domain.userService interface
func NewUserUseCase(c *UUCConfig) domain.UserUseCase {
	return &userUseCase{
		UserRepository: c.UserRepository,
	}
}

func (uuc *userUseCase) Get(ctx context.Context, uid uuid.UUID) (*domain.User, error) {
	u, err := uuc.UserRepository.FindByID(ctx, uid)

	return u, err
}

func (s *userUseCase) Login(ctx context.Context, user *domain.User) error {

	uFetched, err := s.UserRepository.FindByEmail(ctx, user.Email)

	// Will return NotAuthorized to client to omit details of why
	if err != nil {
		return apperrors.NewAuthorization("Invalid email and password combination")
	}
	// verify password - we previously created this method
	match, err := comparePasswords(uFetched.Password, user.Password)

	if err != nil {
		return apperrors.NewInternal()
	}

	if !match {
		return apperrors.NewAuthorization("Invalid email and password combination")
	}

	*user = *uFetched
	return nil
}
