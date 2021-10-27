package author

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset    int   `json:"offset"`
	Forward   bool  `json:"forward"`
	MessageId int64 `json:"messageId"`
}

func (c *LicenseAuthorCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("LicenseAuthorCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	tempMsg := tgbotapi.Message{}
	tempChat := tgbotapi.Chat{}
	tempMsg.Chat = &tempChat
	tempMsg.Chat.ID = parsedData.MessageId

	if parsedData.Forward {
		c.List(uint64(parsedData.Offset+limit), &tempMsg)
	} else {
		position := parsedData.Offset - limit
		if position < 0 {
			position = 0
		}
		c.List(uint64(position), &tempMsg)
	}

}
