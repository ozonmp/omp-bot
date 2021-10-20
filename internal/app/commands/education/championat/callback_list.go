package championat

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"Offset"`
}

func (c *ChampionatCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("ChampionatCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	const defaultLimit = 3
	outputMsgText := fmt.Sprintf("List of the championat (from %v to %v)\n\n", parsedData.Offset, parsedData.Offset+defaultLimit-1)
	championats := c.championatService.List(parsedData.Offset, defaultLimit)
	for _, p := range championats {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	if len(championats) == 0 {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "No more championat!")
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("ChampionatCommander.CallbackList: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: parsedData.Offset + defaultLimit,
	})

	newCallbackPath := path.CallbackPath{
		Domain:       "education",
		Subdomain:    "championat",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", newCallbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ChampionatCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
