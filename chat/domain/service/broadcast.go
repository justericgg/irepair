package service

import "github.com/justericgg/irepair/chat/domain/model/room"

type BroadcastSvc interface {
	Broadcast(room.Room) error
}
