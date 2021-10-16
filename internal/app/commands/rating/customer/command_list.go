package customer

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CustomerCommander) List(inputMessage *tgbotapi.Message) error {
	outputMsgText := "Here all the customers: \n\n"

	customers, err := c.customerService.List(0, 10000)
	if err != nil {
		return err
	}
	for i, p := range customers {
		outputMsgText += fmt.Sprintf("%d. %s", i, p.Title)
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, err := json.Marshal(CallbackListData{
		Offset: 0,
	})
	if err != nil {
		return err
	}

	callbackPath := path.CallbackPath{
		Domain:       "rating",
		Subdomain:    "customer",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	return err
}
