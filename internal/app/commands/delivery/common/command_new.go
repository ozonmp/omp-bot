package common

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/delivery"
	"log"
)

func (c DummyCommonCommander) New(inputMsg *tgbotapi.Message) {
	jsonString := inputMsg.CommandArguments()
	if jsonString == "" {
		log.Println("Empty arg, expected JSON string")
		return
	}

	var commons []delivery.Common
	err := json.Unmarshal([]byte(jsonString), &commons)
	if err != nil {
		log.Printf("DummyCommonCommander.New: "+
			"error reading json data for type Common from "+
			"input string %v - %v", jsonString, err)
		return
	}

	text := ""
	for _, common := range commons {
		commonId, err := c.commonService.Create(common)
		if err == nil {
			text += fmt.Sprintf("New delivery created with id: %d\n", commonId)
		} else {
			text += fmt.Sprintf("New delivery creation failed: %s\n", err)
		}
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, text)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummyCommonCommander.New: error sending reply message to chat - %v", err)
	}
}
