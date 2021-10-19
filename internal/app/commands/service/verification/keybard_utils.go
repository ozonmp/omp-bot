package verification

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *ServiceVerificationCommander) createKeyboard(position uint64, currItemsAmount uint64) []tgbotapi.InlineKeyboardButton {

	dataLen := c.verificationService.GetDataLen()

	var buttons = make([]tgbotapi.InlineKeyboardButton, 0)

	// Back page button + Home page button
	var lastPosition = int64(position) - int64(currItemsAmount)
	if lastPosition >= 0 {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("⬅️", c.createButton(uint64(lastPosition)).String()),
			tgbotapi.NewInlineKeyboardButtonData("🏠", c.createButton(0).String()))
	}

	// Next page button
	var nextPosition = int64(position) + int64(currItemsAmount)
	if nextPosition <= int64(dataLen-1) {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("➡️", c.createButton(uint64(nextPosition)).String()))
	}

	return buttons
}

func (c *ServiceVerificationCommander) createButton(dataOffset uint64) path.CallbackPath {
	serializedData, _ := json.Marshal(CallbackListData{
		Offset: dataOffset,
	})

	callbackPath := path.CallbackPath{
		Domain:       "service",
		Subdomain:    "verification",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return callbackPath
}