package assets

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *AssetsCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__bank__assets - this help\n" +
			"/list__bank__assets - list assets by page\n" +
			"/get__bank__assets <ID> - get an asset\n" +
			"/delete__bank__assets <ID> - remove existing asset\n\n" +

			"/new__bank__assets <userId> <price> - create new asset\n" +
			"/edit__bank__assets <ID> <userId> <price> - edit existing asset\n",
	)

	_, err := c.Bot.Send(msg)
	if err != nil {
		log.Printf("AssetsCommander.Help: error sending reply message to chat - %v", err)
	}
}
