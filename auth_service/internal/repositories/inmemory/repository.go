package inmemory

import (
	"auth_service/internal/app/models"
	"sync"
)

type storage struct {
	registerUsers map[models.Email]user
	authUsers     map[models.Token]user
}

type Repository struct {
	storage storage
	mx      sync.Mutex
}

func NewRepository(cap int) *Repository {
	return &Repository{
		storage: storage{
			registerUsers: make(map[models.Email]user, cap),
			authUsers:     make(map[models.Token]user, cap),
		},
	}
}

type user struct {
	*models.User
}
