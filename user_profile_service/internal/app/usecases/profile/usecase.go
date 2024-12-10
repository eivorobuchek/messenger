package profile

import (
	"context"
	"user_profile_service/internal/app/models"
)

type UsecaseProfile interface {
	EditUserProfile(ctx context.Context, profile *EditProfile) error
	FindUsersByNickname(ctx context.Context, nickname models.Nickname) ([]models.Profile, error)
}

// RepositoryProfile описывает методы для работы с профилями
type RepositoryProfile interface {
	EditUserProfile(ctx context.Context, user models.Profile) error
	FindUsersByNickname(ctx context.Context, nickname models.Nickname) ([]models.Profile, error)
}

type DepsProfile struct {
	ProfileRepository RepositoryProfile
}

type profileUsecase struct {
	DepsProfile
}

// NewProfileUsecase профиль
func NewProfileUsecase(d DepsProfile) *profileUsecase {
	return &profileUsecase{DepsProfile: d}
}
