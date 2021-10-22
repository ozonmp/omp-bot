package product

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *ProductCommanderImpl) sendPage(chatID int64, page int) {
	products, err := c.service.List(uint64(page-1)*c.pageSize, c.pageSize)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error geting product list - %v", err)
	}

	outputMsgText := ""
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(chatID, outputMsgText)

	if uint64(page)*c.pageSize < c.service.Count() {
		addNextPageButton(&msg, page+1)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}

func addNextPageButton(msg *tgbotapi.MessageConfig, page int) {
	serializedData, _ := json.Marshal(CallbackListData{
		Page: page,
	})

	callbackPath := path.CallbackPath{
		Domain:       "lists",
		Subdomain:    "product",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)
}
