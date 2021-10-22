package referral

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset    int   `json:"o"`
	Forward   bool  `json:"f"`
	MessageId int64 `json:"i"` //data limit -64 bytes https://core.telegram.org/bots/api#inlinekeyboardbutton
}

func (c *ReferralCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	if err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData); err != nil {
		log.Printf("ReferralCommander.CallbackList Cant parse data %s", callbackPath.CallbackData)
		return
	}
	tempMsg := tgbotapi.Message{}
	tempChat := tgbotapi.Chat{}
	tempMsg.Chat = &tempChat
	tempMsg.Chat.ID = parsedData.MessageId // средней элегатности решение

	if parsedData.Forward {
		c.List(uint64(parsedData.Offset+int(pageSize)), &tempMsg)
	} else {
		position := parsedData.Offset - int(pageSize)
		if position < 0 {
			position = 0
		}
		c.List(uint64(position), &tempMsg)
	}
}
