package exchange

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *SubdomainCommander) List(inputMsg *tgbotapi.Message) {
	c.ShowPage(inputMsg, 0)
}

func (c *SubdomainCommander) ShowPage(inputMsg *tgbotapi.Message, cursor uint64) {
	outputMsgText := "Here all the exchange requests: \n\n"

	exchangeRequestList, err := c.exchangeService.List(cursor, 5)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
			"You've reached end of list",
		)

		_, err2 := c.bot.Send(msg)
		if err2 != nil {
			log.Printf("SubdomainCommander.List: error sending reply message to chat - %v", err2)
		}
		return
	}
	for _, er := range exchangeRequestList {
		outputMsgText += fmt.Sprintf("ID: %v, Package: %v\nFrom: %v, To: %v\nStatus: %v\n",
			er.Id, er.Package, er.From, er.To, er.Status)
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: cursor + 5,
	})

	callbackPath := path.CallbackPath{
		Domain:       "exchange",
		Subdomain:    "exchange",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}