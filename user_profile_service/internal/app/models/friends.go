package models

type Friends map[UserId]Profile

type FriendRequest struct {
	RequesterID UserId
	ReceiverID  UserId
	Status      string // "pending", "accepted", "rejected"
}

func (f *Friends) CheckFriend(friendId UserId) bool {
	if _, ok := (*f)[friendId]; ok {
		return true
	}
	return false
}
