package visit

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *VisitCommanderStruct) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	visitId, err := strconv.Atoi(args)
	if err != nil {
		c.Send(inputMessage.Chat.ID, "Please, specify visit id")
		log.Println("wrong args", args)
		return
	}

	visit, err := c.visitService.Describe(uint64(visitId))
	if err != nil {
		c.Send(inputMessage.Chat.ID, "Visit not found")
		log.Printf("fail to get visit with id %d: %v", visitId, err)
		return
	}

	c.Send(inputMessage.Chat.ID, visit.Title)
}
