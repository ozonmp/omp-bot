package ground

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	Limit        = 2
	PrevPageText = "Prev page"
	NextPageText = "Next Page"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *GroundCommander) List(inputMessage *tgbotapi.Message) {
	msgText := "All grounds: \n\n"

	grounds, err := c.service.List(0, Limit)
	if err != nil {
		log.Printf("Internal error %v", err)
		c.Send(inputMessage.Chat.ID, "Failed to get a list of grounds")
		return
	}

	for i, g := range grounds {
		msgText += fmt.Sprintf("%d. %s", i, g.String())
		msgText += "\n"
	}

	if c.service.Count() > Limit {
		serializedData, _ := json.Marshal(CallbackListData{
			Cursor: Limit,
			Limit:  Limit,
		})

		callbackPath := path.CallbackPath{
			Domain:       "autotransport",
			Subdomain:    "ground",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(NextPageText, callbackPath.String()),
			),
		)

		c.SendWithReply(inputMessage.Chat.ID, msgText, replyMarkup)
		msgText = ""
	}
}
