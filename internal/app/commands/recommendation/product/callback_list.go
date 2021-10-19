package product

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (commander *ProductCommander) ShowButtons(cursor uint64, limit uint64, chatId int64) {
	size := commander.service.Size()
	if size == 0 {
		commander.Send(chatId, "Product list is empty")
		return
	}
	products, err := commander.service.List(cursor, limit)
	if err != nil {
		commander.Send(chatId, err.Error())
		return
	}

	msgText := ""

	for _, p := range products {
		msgText += p.String()
		msgText += "\n"
	}

	msg := tgbotapi.NewMessage(chatId, msgText)

	serializedData, error := json.Marshal(CallbackListData{
		Cursor: cursor,
		Limit:  limit,
	})

	if error != nil {
		commander.Send(chatId, error.Error())
		return
	}

	var nextButton tgbotapi.InlineKeyboardButton
	var prevButton tgbotapi.InlineKeyboardButton
	if cursor+limit < uint64(size) {
		callbackNextPagePath := path.CallbackPath{
			Domain:       "recommendation",
			Subdomain:    "product",
			CallbackName: "list_next",
			CallbackData: string(serializedData),
		}
		nextButton = tgbotapi.NewInlineKeyboardButtonData("Next page", callbackNextPagePath.String())
	}
	if cursor > 0 {
		callbackPrevPagePath := path.CallbackPath{
			Domain:       "recommendation",
			Subdomain:    "product",
			CallbackName: "list_prev",
			CallbackData: string(serializedData),
		}
		prevButton = tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPrevPagePath.String())
	}
	switch {
	case len(nextButton.Text) != 0 && len(prevButton.Text) != 0:
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				prevButton, nextButton,
			),
		)

	case len(prevButton.Text) != 0:
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				prevButton,
			),
		)

	case len(nextButton.Text) != 0:
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				nextButton,
			),
		)

	}

	if _, err := commander.bot.Send(msg); err != nil {
		log.Println(err)
	}

}

func (commander *ProductCommander) CallbackNextList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil || len(callbackPath.CallbackData) == 0 {
		log.Println("Not correct data")
	}
	cursorStartPos := parsedData.Cursor + parsedData.Limit
	commander.ShowButtons(cursorStartPos, parsedData.Limit, callback.Message.Chat.ID)
}

func (commander *ProductCommander) CallbackPrevList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil || len(callbackPath.CallbackData) == 0 {
		log.Println("Not correct data")
	}
	cursorStartPos := parsedData.Cursor - parsedData.Limit
	commander.ShowButtons(cursorStartPos, parsedData.Limit, callback.Message.Chat.ID)
}
