package friendship

import (
	"user_profile_service/internal/app/usecases/friendsfip"
)

type Deps struct {
	FriendshipUsecase friendsfip.FriendshipUsecase
}

type Controller struct {
	Deps
}

// NewFriendshipHandler создает новый экземпляр обработчика профиля
func NewFriendshipHandler(deps Deps) *Controller {
	return &Controller{Deps: deps}
}
