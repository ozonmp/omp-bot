package incident

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	entity, err := c.incidentService.Get(idx)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.MessageConfig{}
	rawEntity, err := entity.String()
	if err != nil {
		log.Printf("Convert incident to string error: %v", err)
		return
	} else {
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			rawEntity,
		)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingIncidentCommander.Get: error sending reply message to chat - %v", err)
	}
}
