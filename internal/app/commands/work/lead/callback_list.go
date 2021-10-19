package lead

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type callbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *LeadCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := callbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("Work.LeadCommander.CallbackList: "+
			"error reading json data for type callbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	items, _ := c.leadService.List(parsedData.Offset, defaultListLimit)

	c.sendListMsg(callback.Message.Chat.ID, paginatedList{
		Items:  items,
		Offset: parsedData.Offset,
		Limit:  defaultListLimit,
	})
}
