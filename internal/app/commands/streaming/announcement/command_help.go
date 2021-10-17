package announcement

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *StreamingAnnouncementCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__streaming__announcement - help\n"+
		"/list__streaming__announcement - list announcements\n"+
		"/get__streaming__announcement - get announcement info\n"+
		"/edit__streaming__announcement - edit announcement info\n"+
		"/new__streaming__announcement - add new announcement\n"+
		"/delete__streaming__announcement - delete announcement",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.Help: error sending reply message to chat - %v", err)
	}
}