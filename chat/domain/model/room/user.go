package room

type User struct {
	connectionId string
}

func CreateUser(connId string) User {
	return User{connectionId: connId}
}
