package serial

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaSerialCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__cinema__serial - help\n"+
			"/list__cinema__serial - list serials\n"+
			"/get__cinema__serial <id> -  get serial\n"+
			"/new__cinema__serial - new serial\n"+
			"/edit__cinema__serial <id> - edit serial\n"+
			"/delete__cinema__serial <id> - delete serial\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CinemaSerialCommander.Help: error sending reply message to chat (%v)", err)
	}
}
