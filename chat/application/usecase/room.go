package usecase

import (
	"github.com/justericgg/irepair/chat/domain/model/room"
	"github.com/justericgg/irepair/chat/domain/repository"
)

type JoinRoomSvc struct {
	repository repository.RoomRepository
}

func NewJoinRoomSvc(repo repository.RoomRepository) *JoinRoomSvc {
	return &JoinRoomSvc{repo}
}

func (svc *JoinRoomSvc) Join(connectionId string) error {

	user := room.CreateUser(connectionId)
	theRoom := room.NewRoom([]room.User{user})

	err := svc.repository.Save(*theRoom)
	if err != nil {
		return err
	}

	return nil
}
