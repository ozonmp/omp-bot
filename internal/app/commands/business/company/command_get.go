package company

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CompanyCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("CompanyCommander.Get: wrong args", args)
		return
	}

	company, err := c.companyService.Describe(uint64(idx))
	if err != nil {
		log.Printf("CompanyCommander.Get: fail to get company with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		company.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.Get: error sending reply message to chat - %v", err)
	}
}
