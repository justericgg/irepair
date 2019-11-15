package service

import (
	"github.com/justericgg/irepair/chat/domain/model/room"
	"time"
)

type Bot struct{}

func (b *Bot) AddMessageIfNeed(r *room.Room, message string) {

	switch message {
	case "誰最漂亮":
		m := room.Message{
			Id:      "BOT001",
			Author:  "魔鏡",
			Message: "是9N唷~",
			Time:    int(time.Now().Unix()),
		}
		r.ReceiveMessage(m)
	case "我誰":
		m := room.Message{
			Id:      "BOT002",
			Author:  "蛇丸寶貝",
			Message: "我瘋子~",
			Time:    int(time.Now().Unix()),
		}
		r.ReceiveMessage(m)
	}
}
