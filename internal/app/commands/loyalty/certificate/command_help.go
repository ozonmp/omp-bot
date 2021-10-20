package certificate

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LoyaltyCertificateCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__{domain}__{subdomain} - help\n"+
			"/list__{domain}__{subdomain} - list of entities\n"+
			"/get__{domain}__{subdomain} - get entity by index\n"+
			"/delete__{domain}__{subdomain} - delete an entity\n"+
			"/new__{domain}__{subdomain} — create a new entity\n"+
			"/edit__{domain}__{subdomain} — edit a entity",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.Help: error sending reply message to chat - %v", err)
	}
}

