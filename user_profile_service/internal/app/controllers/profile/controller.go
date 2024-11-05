package profile

import (
	usecase "user_profile_service/internal/app/usecases/profile"
)

type DepsProfile struct {
	UsecaseProfile *usecase.UsecaseProfile
}

type ControllerProfile struct {
	DepsProfile
}

// NewProfileHandler создает новый экземпляр обработчика профиля
func NewProfileHandler(deps DepsProfile) *ControllerProfile {
	return &ControllerProfile{
		DepsProfile: deps,
	}
}
