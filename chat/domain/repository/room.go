package repository

import "github.com/justericgg/irepair/chat/domain/model/room"

type RoomRepository interface {
	Save(room.Room) error
	Delete(room.Room) error
	BuildRoomWithUsers() (*room.Room, error)
}
