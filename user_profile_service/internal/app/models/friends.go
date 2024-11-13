package models

import "time"

type Friends map[UserId]Profile

type FriendRequest struct {
	RequesterID UserId
	ReceiverID  UserId
	Status      string // "pending", "accepted", "rejected"
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (f *Friends) CheckFriend(friendId UserId) bool {
	if _, ok := (*f)[friendId]; ok {
		return true
	}
	return false
}
