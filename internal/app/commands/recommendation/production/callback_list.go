package production

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	MessageId int
	Offset    int
}

func (c *RecommendationProductionCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var data CallbackListData
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &data)
	if err != nil {
		log.Printf("%v parse callback error %v", commanderName, err)
		return
	}

	data.Offset += listLimit

	c.list(callback.Message.Chat.ID, data)
}
