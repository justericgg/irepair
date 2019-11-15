package service

import "github.com/justericgg/irepair/chat/domain/model/room"

type Bot interface {
	Response(room.Room) error
}
