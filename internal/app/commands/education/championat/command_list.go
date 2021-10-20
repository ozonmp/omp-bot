package championat

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *ChampionatCommander) List(inputMessage *tgbotapi.Message) {
	const defaultLimit = 3
	outputMsgText := fmt.Sprintf("List of the championat (from %v to %v)\n\n", 0, defaultLimit-1)
	championats := c.championatService.List(0, defaultLimit)
	for _, p := range championats {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: defaultLimit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "education",
		Subdomain:    "championat",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ChampionatCommander.List: error sending reply message to chat - %v", err)
	}
}
