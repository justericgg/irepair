package room

type Room struct {
	users []User
}

func NewRoom(users []User) *Room {
	return &Room{users}
}

func (r *Room) GetFirstUserConnId() string {
	return r.users[0].connectionId
}
