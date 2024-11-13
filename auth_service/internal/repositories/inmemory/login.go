package inmemory

import (
	"auth_service/internal/app/models"
	"context"
	"fmt"
)

func (r *Repository) Login(ctx context.Context, userInfo *models.User, token models.Token) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	_, ok := r.storage.registerUsers[userInfo.Email]
	if !ok {
		return fmt.Errorf("account with email '%s': %w", userInfo.Email, models.ErrNotFound)
	}

	r.storage.authUsers[token] = user{userInfo}
	return nil
}
