package order

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *OrderCommander) List(inputMessage *tgbotapi.Message) {
	c.SendPage(inputMessage.Chat.ID, "Here all the orders: \n\n", 0, 5)
}

func (c *OrderCommander) SendPage(chatID int64, header string, cursor uint64, limit uint64) {
	orders, err := c.orderService.List(cursor, limit+1)
	if err != nil {
		log.Printf("OrderCommander.SendPage: error getting order list - %v", err)
		return
	}

	msg := tgbotapi.NewMessage(chatID, "")

	if len(orders) == int(limit+1) {
		addNextButton(&msg, cursor, limit)
		orders = orders[:len(orders)-1]
	}

	outputMsgText := header
	for _, p := range orders {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg.Text = outputMsgText

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("OrderCommander.SendPage: error sending reply message to chat - %v", err)
	}
}

func addNextButton(msg *tgbotapi.MessageConfig, cursor uint64, limit uint64) {
	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: cursor + limit,
		Limit:  limit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "buy",
		Subdomain:    "order",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)
}
