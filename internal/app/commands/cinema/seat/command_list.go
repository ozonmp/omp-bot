package seat

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CinemaSeatCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the seats: \n\n"

	seats, _ := c.subdomainService.List(1, 10)
	for _, seat := range seats {
		outputMsgText += fmt.Sprintf("Seat #%d (Row: %d Number %d) price: %d", seat.ID, seat.Row, seat.Number, seat.Price)
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "cinema",
		Subdomain:    "seat",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	c.bot.Send(msg)
}
