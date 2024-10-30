package usecases

import (
	"auth_service/internal/app/models"
	"context"
	"errors"
	"fmt"
)

var (
	ErrInvalidValue    = errors.New("invalid value")
	RegisterUserFailed = errors.New("register user failed")
	UnregisterUser     = errors.New("unregister user")
)

func (uc *authUsecase) RegisterUser(ctx context.Context, userInfo *RegisterUser) (*models.User, error) {
	// Валидация
	if userInfo == nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidValue, "empty request")
	}
	err := userInfo.Validate()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	// Создание пользователя
	user := models.NewUser()
	user.Email = userInfo.Email
	user.HashedPassword, err = models.HashPassword(userInfo.Password)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	// Сохранение в БД
	if err := uc.AuthRepository.RegisterUser(ctx, user); err != nil {
		return nil, fmt.Errorf("%w: %s", RegisterUserFailed, err.Error())
	}

	return user, nil
}

func (uc *authUsecase) Login(ctx context.Context, loginUser *LoginUser) (*models.AuthUser, error) {
	// Валидация
	if loginUser == nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidValue, "empty request")
	}
	err := loginUser.Validate()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	accountUser, err := uc.AuthRepository.GetRegisterUserByEmail(ctx, models.Email(loginUser.Email))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", UnregisterUser, err.Error())
	}

	authUser := models.NewAuthUser()
	authUser.User = accountUser
	token, err := models.NewToken(accountUser)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", RegisterUserFailed, err.Error())
	}
	authUser.Token = token

	return authUser, nil
}
