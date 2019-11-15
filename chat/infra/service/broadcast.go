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

const Action = "sendmessage"

type Payload struct {
	Action string       `json:"action"`
	Data   room.Message `json:"data"`
}

func (svc *BroadcastSvc) Broadcast(theRoom *room.Room) error {

	conn, err := adapter.GetConnection()
	if err != nil {
		return err
	}

	for _, connId := range theRoom.GetAllUserConnId() {

		for _, message := range theRoom.GetMessages() {

			payload := Payload{
				Action: Action,
				Data:   message,
			}

			data, err := json.Marshal(payload)
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
