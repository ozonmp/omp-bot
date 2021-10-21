package operation

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit uint64 `json:"limit"`
}

func (c *BankOperationCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("BankOperationCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	cursor, limit := parsedData.Cursor, parsedData.Limit

	msg, err := c.prepareList(*callback.Message, cursor, limit)

	if err != nil {
		log.Printf("BankOperationCommander.List: error preparing list - %v", err)
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankOperationCommander.List: error sending reply message to chat - %v", err)
	}
}