package usecase

import (
	"encoding/json"
	"github.com/justericgg/irepair/chat/domain/model/room"
	"github.com/justericgg/irepair/chat/domain/repository"
	"github.com/justericgg/irepair/chat/infra/service"
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

type LeaveRoomSvc struct {
	repository repository.RoomRepository
}

func NewLeaveRoomSvc(repo repository.RoomRepository) *LeaveRoomSvc {
	return &LeaveRoomSvc{repo}
}

func (svc *LeaveRoomSvc) Leave(connectionId string) error {

	user := room.CreateUser(connectionId)
	theRoom := room.NewRoom([]room.User{user})

	err := svc.repository.Delete(*theRoom)
	if err != nil {
		return err
	}

	return nil
}

type MessageSvc struct {
	repository   repository.RoomRepository
	broadcastSvc service.BroadcastSvc
}

func NewMessageSvc(rRepo repository.RoomRepository, svc service.BroadcastSvc) *MessageSvc {
	return &MessageSvc{rRepo, svc}
}

func (svc *MessageSvc) ProcessMessage(endpoint, request string) error {

	message := &room.Message{}
	err := json.Unmarshal([]byte(request), &message)
	if err != nil {
		return err
	}

	theRoom, err := svc.repository.BuildRoomWithUsers()
	if err != nil {
		return err
	}
	theRoom.ReceiveMessage(*message)

	svc.broadcastSvc.Endpoint = endpoint
	err = svc.broadcastSvc.Broadcast(*theRoom)
	if err != nil {
		return err
	}

	return nil
}
