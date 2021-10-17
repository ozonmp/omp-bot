package subdomain

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummySubdomainCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("DummySubdomainCommander.Get invalid args", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageGet))
		return
	}

	elem, err := c.subdomainService.Describe(uint64(idx))
	if err != nil {
		log.Printf("DummySubdomainCommander.Get fail to describe elem with idx %d: %v\n", idx, err)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, ErrNotFound))
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		elem.String(),
	)

	c.bot.Send(msg)
}
