package frontend

import (
	"context"

	"github.com/aftosmiros/moodu/user-service/internal/domain"
)

type UserUsecase interface {
	Register(ctx context.Context, user *domain.User, password string) (id string, err error)
	Login(ctx context.Context, email, password string) (*domain.Token, error)
	GetProfile(ctx context.Context, userID string) (*domain.User, error)
	RefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error)
	Logout(ctx context.Context, refreshToken string) error
	UpdateProfile(ctx context.Context, input domain.UpdateProfileInput) error
	ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error
}
