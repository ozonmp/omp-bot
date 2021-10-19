package incident

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	msg := tgbotapi.MessageConfig{}

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"wrong format id",
		)
	} else {
		err = c.incidentService.Delete(idx)
		if err != nil {
			log.Printf("fail to edit incident with id %d: %v", idx, err)
			msg = tgbotapi.NewMessage(
				inputMessage.Chat.ID,
				"incident not found",
			)
		} else {
			msg = tgbotapi.NewMessage(
				inputMessage.Chat.ID,
				"done",
			)
		}
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingIncidentCommander.Delete: error sending reply message to chat - %v", err)
	}
}
