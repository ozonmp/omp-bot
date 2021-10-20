package seat

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *CinemaSeatCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	id, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Printf("Delete: wrong arguments: %s", args)

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"wrong args",
		)
		c.bot.Send(msg)

		return
	}

	_, err = c.subdomainService.Remove(id)
	if err != nil {
		log.Printf("Delete: error seat: %v", err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Error",
		)
		c.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Success deleted",
	)
	c.bot.Send(msg)
}
