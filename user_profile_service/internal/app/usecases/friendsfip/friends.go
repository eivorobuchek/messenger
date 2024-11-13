package friendsfip

import (
	"context"
	"errors"
	"fmt"
	"user_profile_service/internal/app/models"
)

var (
	ErrInvalidValue = errors.New("invalid value")
	ErrRepo         = errors.New("repo error")
)

// AddFriend добавляет пользователя в друзья
func (u *friendshipUsecase) AddFriend(ctx context.Context, friendRequest *AddFriendRequest) error {
	if friendRequest == nil {
		return fmt.Errorf("%w: %s", ErrInvalidValue, "empty request")
	}
	err := friendRequest.Validate()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}
	return u.FriendshipRepository.AddFriendRequest(ctx, friendRequest.FriendId, friendRequest.UserId)
}

// RemoveFriend удаляет пользователя из друзей
func (u *friendshipUsecase) RemoveFriend(ctx context.Context, deleteFriend *DeleteFriendRequest) error {
	if deleteFriend == nil {
		return fmt.Errorf("%w: %s", ErrInvalidValue, "empty request")
	}
	err := deleteFriend.Validate()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}
	return u.FriendshipRepository.RemoveFriendRequest(ctx, deleteFriend.FriendId, deleteFriend.UserId)
}

// UpdateFriendRequestStatus обновляет статус запроса на дружбу
func (u *friendshipUsecase) UpdateFriendRequestStatus(ctx context.Context, updateFriendship *UpdateFriendshipStatus) error {

	if updateFriendship == nil {
		return fmt.Errorf("%w: %s", ErrInvalidValue, "empty request")
	}
	err := updateFriendship.Validate()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	return u.FriendshipRepository.UpdateFriendRequestStatus(ctx, updateFriendship.FriendId, updateFriendship.UserId, updateFriendship.Status)
}

// GetFriends возвращает список друзей пользователя
func (u *friendshipUsecase) GetFriends(ctx context.Context, findFriends *GetFriendsRequest) ([]models.Profile, error) {
	if findFriends == nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidValue, "empty request")
	}
	err := findFriends.Validate()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}
	friends, err := u.FriendshipRepository.GetFriends(ctx, findFriends.UserId)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrRepo, err.Error())
	}
	return friends, nil
}
