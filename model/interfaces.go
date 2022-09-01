package model

import (
	"context"
	"time"
)

type UserService interface {
	GetById(ctx context.Context, id uint) (*User, error)
	SignUp(ctx context.Context, user *User) error
	SignIn(ctx context.Context, user *User) error
}

type UserRepository interface {
	FindById(ctx context.Context, id uint) (*User, error)
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type TokenService interface {
	GetNewPairFromUser(ctx context.Context, user *User, prevTokenId string) (*TokenPair, error)
	ValidateIdToken(tokenString string) (*User, error)
	ValidateRefreshToken(refresh_token string) (*RefreshToken, error)
}

type TokenRepository interface {
	SetNewRefreshToken(ctx context.Context, userId string, tokenId string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userId string, prevTokenId string) error
}
