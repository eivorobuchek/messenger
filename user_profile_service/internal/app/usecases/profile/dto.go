package profile

import (
	"errors"
	"user_profile_service/internal/app/models"
)

type EditProfile struct {
	Nickname        models.Nickname
	ChangedNickname models.Nickname
	Bio             string
	Avatar          string
}

func (p EditProfile) Validate() error {
	if p.Nickname.String() == "" {
		return errors.New("nickname is required")
	}

	return nil
}
