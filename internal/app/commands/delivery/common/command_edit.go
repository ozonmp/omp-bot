package common

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/delivery"
	"log"
)

type newCommon struct {
	Id        uint64
	NewEntity delivery.Common
}

func (c DummyCommonCommander) Edit(inputMsg *tgbotapi.Message) {
	jsonString := inputMsg.CommandArguments()
	if jsonString == "" {
		log.Println("Empty ARG, but expected JSON string")
		return
	}

	var newCommons []newCommon
	err := json.Unmarshal([]byte(jsonString), &newCommons)
	if err != nil {
		log.Printf("DummyCommonCommander.Edit: "+
			"error reading json data for type newCommon from "+
			"input string %v - %v", jsonString, err)
		return
	}

	text := ""
	for _, newCommon := range newCommons {
		err := c.commonService.Update(newCommon.Id, newCommon.NewEntity)
		if err == nil {
			text += fmt.Sprintf("Delivery with id: %d was updated\n", newCommon.Id)
		} else {
			text += fmt.Sprintf("New delivery creation failed: %s\n", err)
		}
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, text)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummyCommonCommander.Edit: error sending reply message to chat - %v", err)
	}
}
