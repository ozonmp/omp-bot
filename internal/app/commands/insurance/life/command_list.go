package life

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/insurance"
)

const limit = 3
const header = "Lifes list: \n\n"

type CallbackListData struct {
	Pointer uint64 `json:"pointer"`
}

func (telegramLifeCommander *TelegramLifeCommander) List(inputMessage *tgbotapi.Message) {

	lifes, err := telegramLifeCommander.lifeService.List(0, limit)
	outputMsgText := header + lifesToString(lifes...)

	message := tgbotapi.NewMessage(inputMessage.Chat.ID, header+outputMsgText)

	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			getButtons(0, err)...,
		),
	)

	telegramLifeCommander.bot.Send(message)

}

func (telegramLifeCommander *TelegramLifeCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)

	lifes, err := telegramLifeCommander.lifeService.List(parsedData.Pointer, limit)

	outputMsgText := header + lifesToString(lifes...)

	message := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	message.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			getButtons(parsedData.Pointer, err)...,
		),
	)

	telegramLifeCommander.bot.DeleteMessage(tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID))
	telegramLifeCommander.bot.Send(message)
}

func getButtons(pointer uint64, err error) []tgbotapi.InlineKeyboardButton {

	var buttons []tgbotapi.InlineKeyboardButton

	if pointer > 0 {
		buttons = append(buttons,
			tgbotapi.NewInlineKeyboardButtonData("Previous page", getCallBackPath(pointer-limit).String()))
	}

	if err == nil {
		buttons = append(buttons,
			tgbotapi.NewInlineKeyboardButtonData("Next page", getCallBackPath(pointer+limit).String()))
	}

	return buttons
}

func getCallBackPath(pointer uint64) path.CallbackPath {
	serializedData, _ := json.Marshal(CallbackListData{
		Pointer: pointer,
	})

	newCallbackPath := path.CallbackPath{
		Domain:       "insurance",
		Subdomain:    "life",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return newCallbackPath
}

func lifesToString(lifes ...insurance.Life) string {
	outputMsgText := ""
	for _, life := range lifes {
		lifeJson, _ := json.Marshal(life)
		outputMsgText += string(lifeJson)
		outputMsgText += "\n"
	}
	return outputMsgText
}
