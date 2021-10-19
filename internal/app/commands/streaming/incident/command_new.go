package incident

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/streaming/incident"
)

func (c *Commander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	msg := tgbotapi.MessageConfig{}

	entity, err := incident.ConvertStringToIncident(args)
	if err != nil {
		log.Println("wrong args", args)
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"wrong format incident",
		)
	} else {
		err = c.incidentService.New(entity)
		if err != nil {
			log.Printf("fail to create incident with id %d: %v", entity.Id, err)
			msg = tgbotapi.NewMessage(
				inputMessage.Chat.ID,
				"incident already exists",
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
		log.Printf("StreamingIncidentCommander.New: error sending reply message to chat - %v", err)
	}
}
