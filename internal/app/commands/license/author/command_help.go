package author

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LicenseAuthorCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		`/help 		- help 
/list__license__author		- list of authors
/get__license__author id 		- get author by id
/delete__license__author id 	- delete author by id
/edit__license__author id firstName lastName 	- edit author info by id
/new__license__author	firstName lastName - add new author`,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LicenseAuthorCommander.Help: error sending reply message to chat - %v", err)
	}
}
