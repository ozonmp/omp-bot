package platform

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (h *PlatformBaseCallbackHandler) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	var msg tgbotapi.MessageConfig

	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("EducationPlatformCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	platforms, err := h.service.List(parsedData.Cursor, parsedData.Limit)
	if err != nil {
		log.Printf(err.Error())

		return
	}

	outputMsgText := ""

	for _, p := range platforms {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	if outputMsgText == "" {
		msg = tgbotapi.NewMessage(callback.Message.Chat.ID, "No more items")
	} else {
		msg = tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
	}

	buttons := makeButtons(parsedData, uint64(len(platforms)))
	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(buttons...),
		)
	}

	_, err = h.bot.Send(msg)
	if err != nil {
		log.Printf("PlatformCallbackHandler: error sending reply message to chat - %v", err)
	}
}

func makeButtons(parsedData CallbackListData, count uint64) []tgbotapi.InlineKeyboardButton {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0, 2)

	if parsedData.Cursor > 0 {
		prevCursor := parsedData.Cursor - DefaultListLimit

		if prevCursor < 0 {
			prevCursor = 0
		}

		serializedData, err := json.Marshal(
			CallbackListData{
				Cursor: prevCursor,
				Limit:  DefaultListLimit,
			},
		)

		if err != nil {
			log.Printf(err.Error())

			serializedData = []byte("")
		}

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(PrevButtonText, getButtonData(serializedData)))
	}

	if count == DefaultListLimit {
		nextCursor := parsedData.Cursor + parsedData.Limit

		serializedData, err := json.Marshal(
			CallbackListData{
				Cursor: nextCursor,
				Limit:  DefaultListLimit,
			},
		)

		if err != nil {
			log.Printf(err.Error())

			serializedData = []byte("")
		}

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(NextButtonText, getButtonData(serializedData)))
	}

	return buttons
}

func getButtonData(serializedData []byte) string {
	callbackPath := path.CallbackPath{
		Domain:       "education",
		Subdomain:    "platform",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return callbackPath.String()
}
