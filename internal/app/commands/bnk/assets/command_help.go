package assets

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *AssetsCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__bnk__assets - this help\n" +
			"/list__bnk__assets - list assets\n" +
			"/get__bnk__assets - get an asset\n" +
			"/delete__bnk__assets - remove existing asset\n\n" +

			"/new__bnk__assets <name>|<money> - create new asset\n" +
			"/edit__bnk__assets <id>|<name>|<money> - edit existing asset\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("AssetsCommander.Help: error sending reply message to chat - %v", err)
	}
}
