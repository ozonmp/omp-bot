package ground

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *GroundCommander) List(inputMessage *tgbotapi.Message)  {
	outputMsgText := "Here all the grounds: \n\n"

	grounds, err := c.service.List(0, 10000)
	if err != nil {
		// return err
	}
	for i, p := range grounds {
		outputMsgText += fmt.Sprintf("%d. %s", i, p.String())
		outputMsgText += "\n"
	}

	serializedData, err := json.Marshal(CallbackListData{
		Offset: 0,
	})
	if err != nil {
		// return err
	}

	callbackPath := path.CallbackPath{
		Domain:       "autotransport",
		Subdomain:    "ground",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	// c.Send(msg)
	c.SendWithReply(inputMessage.Chat.ID, outputMsgText, replyMarkup)
	// return err
}
