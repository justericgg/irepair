package service

import (
	"encoding/json"
	"github.com/justericgg/irepair/chat/domain/model/room"
	"github.com/justericgg/irepair/chat/infra/adapter"
	"log"
)

type BroadcastSvc struct {
	Endpoint string
}

func (svc *BroadcastSvc) Broadcast(theRoom room.Room) error {

	conn, err := adapter.GetConnection()
	if err != nil {
		return err
	}

	for _, connId := range theRoom.GetAllUserConnId() {

		for _, message := range theRoom.GetMessages() {
			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err.Error())
			}
			_, err = conn.Post(svc.Endpoint, connId, data)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
