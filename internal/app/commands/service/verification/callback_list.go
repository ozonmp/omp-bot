package verification

import (
	"encoding/json"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *ServiceVerificationCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	if err != nil {
		log.Printf("ServiceVerificationCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	err = json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	if err != nil {
		log.Printf("ServiceVerificationCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	c.List(parsedData.Offset, callback.Message)
}
