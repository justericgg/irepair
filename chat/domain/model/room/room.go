package room

type Room struct {
	users    []User
	messages []Message
}

func NewRoom(users []User) *Room {
	return &Room{users: users}
}

func (r *Room) GetFirstUserConnId() string {
	return r.users[0].connectionId
}

func (r *Room) GetAllUserConnId() []string {

	connIds := make([]string, len(r.users))
	for _, user := range r.users {
		connIds = append(connIds, user.connectionId)
	}

	return connIds
}

func (r *Room) ReceiveMessage(msg Message) {
	r.messages = append(r.messages, msg)
}

func (r *Room) GetMessages() []Message {
	return r.messages
}
