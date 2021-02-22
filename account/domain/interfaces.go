package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserUseCase interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Login(ctx context.Context, user *User) error
}

type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}

// Token interfaces

type TokenUseCase interface {
	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
	//Signout(ctx context.Context, uid uuid.UUID) error
	//ValidateIDToken(tokenString string) (*User, error)
	//ValidateRefreshToken(refreshTokenString string) (*RefreshToken, error)
}

type TokenRepository interface {
	SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error
	DeleteUserRefreshTokens(ctx context.Context, userID string) error
}
