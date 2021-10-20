package visit

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *VisitCommanderStruct) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	visitId, err := strconv.Atoi(args)
	if err != nil {
		c.Send(inputMessage.Chat.ID, "Visit id is not valid.")
		return
	}

	_, err = c.visitService.Remove(uint64(visitId))
	if err != nil {
		c.Send(inputMessage.Chat.ID, "Can't remove.")
		return
	}

	c.Send(inputMessage.Chat.ID, "Visit removed.")
}
