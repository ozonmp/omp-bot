package rent

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaRentCommander) Help(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	helpString := []string{
		"/help__cinema__rent\tprint list of commands",
		"/get__cinema__rent\tget a entity",
		"/list__cinema__rent\tget a list of your entity",
		"/delete__cinema__rent\tdelete an existing entity",
		"/new__cinema__rent\tcreate a new entity",
		"/edit__cinema__rent\tedit a entity",
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, strings.Join(helpString, "\n"))

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("CinemaRentCommander.Help: Error sending message: %v", err)
	}
}
