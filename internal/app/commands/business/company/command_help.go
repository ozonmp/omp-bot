package company

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CompanyCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__business__company - show this Help\n"+
			"/get__business__company — get a entity\n"+
			"/list__business__company — get a list of your entity\n"+
			"/delete__business__company — delete an existing entity\n\n"+
			"/new__business__company — create a new entity // not implemented\n"+
			"/edit__business__company — create a new entity // not implemented\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CompanyCommander.Help: error sending reply message to chat - %v", err)
	}
}
