package serial

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaSerialCommander) Default(inputMessage *tgbotapi.Message) {
	msg := ""

	defer func() {
		tgmsg := tgbotapi.NewMessage(inputMessage.Chat.ID, msg)
		_, err := c.bot.Send(tgmsg)
		if err != nil {
			log.Printf("CinemaSerialCommander.Default: error sending reply message to chat (%v)", err)
		}
	}()

	msg = fmt.Sprintf("command unknown %s\nUse %s", inputMessage.Text, "/help__cinema__serial")
	log.Printf("CinemaSerialCommander.Default: %s", msg)
}
