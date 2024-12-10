package profile

import (
	"context"
	"errors"
	"fmt"
	"user_profile_service/internal/app/models"
)

var (
	ErrInvalidValue   = errors.New("invalid value")
	ErrGetUserProfile = errors.New("invalid nickname")
)

// EditUserProfile редактирует профиль пользователя
func (p *profileUsecase) EditUserProfile(ctx context.Context, profile *EditProfile) error {
	// Валидация
	if profile == nil {
		return fmt.Errorf("%w: %s", ErrInvalidValue, "empty request")
	}
	err := profile.Validate()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	currentProfile, err := p.ProfileRepository.FindUsersByNickname(ctx, profile.Nickname)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrGetUserProfile, err.Error())
	}

	if len(currentProfile) == 0 {
		return fmt.Errorf("%w: %s", ErrGetUserProfile, profile.Nickname)
	}

	editedProfile := currentProfile[0]
	if profile.ChangedNickname != "" {
		editedProfile.Nickname = profile.ChangedNickname
	}
	if profile.Bio != "" {
		editedProfile.Bio = profile.Bio
	}
	if profile.Avatar != "" {
		editedProfile.Avatar = profile.Avatar
	}

	return p.ProfileRepository.EditUserProfile(ctx, editedProfile)
}

// FindUsersByNickname ищет пользователей по никнейму
func (p *profileUsecase) FindUsersByNickname(ctx context.Context, nickname models.Nickname) ([]models.Profile, error) {
	if nickname == "" {
		return nil, fmt.Errorf("%w: %s", ErrInvalidValue, "empty nickname")
	}

	profile, err := p.ProfileRepository.FindUsersByNickname(ctx, nickname)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrGetUserProfile, err.Error())
	}
	return profile, nil
}
