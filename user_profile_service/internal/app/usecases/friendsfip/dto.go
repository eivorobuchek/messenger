package friendsfip

import (
	"errors"
	"user_profile_service/internal/app/models"
)

type AddFriendRequest struct {
	UserId   models.UserId `json:"UserId"`
	FriendId models.UserId `json:"FriendId"`
}

func (p AddFriendRequest) Validate() error {
	if p.UserId.String() == "" {
		return errors.New("userId is required")
	}
	if p.FriendId.String() == "" {
		return errors.New("friendId is required")
	}

	return nil
}

type DeleteFriendRequest struct {
	UserId   models.UserId `json:"UserId"`
	FriendId models.UserId `json:"FriendId"`
}

func (p DeleteFriendRequest) Validate() error {
	if p.UserId.String() == "" {
		return errors.New("userId is required")
	}
	if p.FriendId.String() == "" {
		return errors.New("friendId is required")
	}

	return nil
}

type UpdateFriendshipStatus struct {
	UserId   models.UserId `json:"UserId"`
	FriendId models.UserId `json:"FriendId"`
	Status   string        `json:"Status"`
}

func (p UpdateFriendshipStatus) Validate() error {
	if p.UserId.String() == "" {
		return errors.New("userId is required")
	}
	if p.FriendId.String() == "" {
		return errors.New("friendId is required")
	}
	if p.Status == "" {
		return errors.New("status is required")
	}

	return nil
}

type GetFriendsRequest struct {
	UserId models.UserId `json:"UserId"`
}

func (p GetFriendsRequest) Validate() error {
	if p.UserId.String() == "" {
		return errors.New("userId is required")
	}
	return nil
}
