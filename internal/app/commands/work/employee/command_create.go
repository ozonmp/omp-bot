package employee

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *DemoSubdomainCommander) Create(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	employee := c.subdomainService.Create(args)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Was created employee["+strconv.Itoa(employee.Id)+"] "+employee.Title,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
