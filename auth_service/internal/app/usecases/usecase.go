package usecases

import (
	"auth_service/internal/app/models"
	"context"
)

type AuthUsecase interface {
	RegisterUser(ctx context.Context, userInfo *RegisterUser) (*models.User, error)
	Login(ctx context.Context, loginUser *LoginUser) (*models.AuthUser, error)
}

type AuthRepository interface {
	RegisterUser(ctx context.Context, o *models.User) error
	Login(ctx context.Context, userInfo *models.User, token models.Token) error
	GetRegisterUserByEmail(ctx context.Context, email models.Email) (*models.User, error)
}

type AuthDeps struct {
	AuthRepository AuthRepository
}

type authUsecase struct {
	AuthDeps
}

// NewAuthUsecase бизнес сервис
func NewAuthUsecase(d AuthDeps) *authUsecase {
	return &authUsecase{
		AuthDeps: d,
	}
}
