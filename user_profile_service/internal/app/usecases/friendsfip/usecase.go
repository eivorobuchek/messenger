package friendsfip

import (
	"context"
	"user_profile_service/internal/app/models"
)

type FriendshipUsecase interface {
	AddFriend(ctx context.Context, friendRequest *AddFriendRequest) error
	RemoveFriend(ctx context.Context, deleteFriend *DeleteFriendRequest) error
	UpdateFriendRequestStatus(ctx context.Context, updateFriendship *UpdateFriendshipStatus) error
	GetFriends(ctx context.Context, findFriends *GetFriendsRequest) ([]models.Profile, error)
}

type FriendshipRepository interface {
	AddFriendRequest(ctx context.Context, friendID, userID models.UserId) error
	RemoveFriendRequest(ctx context.Context, userID, friendID models.UserId) error
	UpdateFriendRequestStatus(ctx context.Context, friendID, userId models.UserId, status string) error
	GetFriends(ctx context.Context, userID models.UserId) ([]models.Profile, error)
}

type FriendshipDeps struct {
	FriendshipRepository FriendshipRepository
}

type friendshipUsecase struct {
	FriendshipDeps
}

// NewFriendshipUsecase профиль
func NewFriendshipUsecase(d FriendshipDeps) *friendshipUsecase {
	return &friendshipUsecase{FriendshipDeps: d}
}
