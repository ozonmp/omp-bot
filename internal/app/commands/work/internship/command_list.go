package internship

import (
	"encoding/json"
	"github.com/VYBel/omp-bot/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *WorkInternshipCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the interships: \n\n"

	products := c.internshipService.List()
	for _, p := range products {
		outputMsgText += c.internshipService.ShortString(p)
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "work",
		Subdomain:    "internship",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternshipCommander.List: error sending reply message to chat - %v", err)
	}
}
