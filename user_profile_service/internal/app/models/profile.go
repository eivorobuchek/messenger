package models

type UserId string

func (id UserId) String() string {
	return string(id)
}

type Nickname string

func (id Nickname) String() string {
	return string(id)
}

type Profile struct {
	UserID   UserId
	Nickname Nickname
	Bio      string
	Avatar   string
	Friends  Friends // IDs друзей
}

func (p *Profile) AddFriend(friend Profile) error {
	if p.Friends.CheckFriend(friend.UserID) {
		return ErrAlreadyFriend
	}

	p.Friends[friend.UserID] = friend

	return nil
}
