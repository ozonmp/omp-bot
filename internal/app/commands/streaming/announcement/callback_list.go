package announcement

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *StreamingAnnouncementCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMsgText := ""
	products, _ := c.announcementService.List(uint64(parsedData.Offset), 5)
	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n----------------------\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	offsetData, _ := json.Marshal(CallbackListData{
		Offset: parsedData.Offset + len(products),
	})
	callbackPath = path.CallbackPath{
		Domain:       "streaming",
		Subdomain:    "announcement",
		CallbackName: "list",
		CallbackData: string(offsetData),
	}
	buttons := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
	}
	if parsedData.Offset > 0 {
		offsetData, _ := json.Marshal(CallbackListData{
			Offset: parsedData.Offset - len(products),
		})
		callbackPath.CallbackData = string(offsetData)
		buttons = append(buttons,
			tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPath.String()),
		)
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(buttons...),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingAnnouncementCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
