package repository

import "github.com/justericgg/irepair/chat/domain/model/room"

type RoomRepository interface {
	Save(room.Room) error
	//delete
}
