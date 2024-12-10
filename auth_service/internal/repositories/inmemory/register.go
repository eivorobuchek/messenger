package inmemory

import (
	"auth_service/internal/app/models"
	"context"
	"fmt"
)

func (r *Repository) RegisterUser(ctx context.Context, o *models.User) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	_, ok := r.storage.registerUsers[o.Email]
	if ok {
		return fmt.Errorf("user email '%s': %w", o.Email, models.ErrAlreadyRegister)
	}

	r.storage.registerUsers[o.Email] = user{o}
	return nil
}

func (r *Repository) GetRegisterUserByEmail(ctx context.Context, email models.Email) (*models.User, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	userInfo, ok := r.storage.registerUsers[email]
	if !ok {
		return nil, fmt.Errorf("user by email '%s': %w", email, models.ErrNotFound)
	}

	return userInfo.User, nil
}
