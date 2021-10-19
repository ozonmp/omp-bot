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
		log.Println("wrong args", args)
		return
	}

	company, err := c.companyService.Get(idx)
	if err != nil {
		log.Printf("fail to get company with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		company.Name+"[zipcode "+strconv.Itoa(int(company.ZipCode))+"]",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.Get: error sending reply message to chat - %v", err)
	}
}