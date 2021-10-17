package subdomain

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *DemoSubdomainCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	products := c.subdomainService.List()
	for _, p := range products {
		outputMsgText += p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, err := json.Marshal(CallbackListData{
		Offset: 21,
	})
	if err != nil {
		log.Println(err)
	}

	callbackPath := path.CallbackPath{
		Domain:       "demo",
		Subdomain:    "subdomain",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}
