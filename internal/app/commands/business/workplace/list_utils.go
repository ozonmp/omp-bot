package workplace

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *BusinessWorkplaceCommander) processList(offset uint64, limit uint64, messageChatId int64) {
	var outputMsgText = fmt.Sprintf("Workplaces(offset - %d, page size - %d, total - %d): \n\n", offset, limit,  c.workplaceService.GetDataSize())

	workplaces, _ := c.workplaceService.List(offset, limit)
	for _, w := range workplaces {
		outputMsgText += w.String()
		outputMsgText += "\n"
	}

	var msg = tgbotapi.NewMessage(messageChatId, outputMsgText)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(c.createTgButtonPanel(offset, uint64(len(workplaces)))...))
	c.bot.Send(msg)
}

func (c *BusinessWorkplaceCommander) createTgButtonPanel(currOffset uint64, currPageSize uint64) []tgbotapi.InlineKeyboardButton {

	var dbDataSize = c.workplaceService.GetDataSize()
	var tgButtons = make([]tgbotapi.InlineKeyboardButton, 0, 4)

	// First page button
	tgButtons = append(tgButtons, tgbotapi.NewInlineKeyboardButtonData("First page", c.createButtonInfo(0).String()))

	// Previous page button
	var prevOffset int64 = int64(currOffset) - int64(currPageSize) - 1
	if prevOffset >= 0 {
		tgButtons = append(tgButtons, tgbotapi.NewInlineKeyboardButtonData("Prev page", c.createButtonInfo(uint64(prevOffset)).String()))
	}

	// Next page button
	var nextOffset = int64(currOffset) + int64(currPageSize)
	if nextOffset <= int64(dbDataSize-1) {
		tgButtons = append(tgButtons, tgbotapi.NewInlineKeyboardButtonData("Next page", c.createButtonInfo(uint64(nextOffset)).String()))
	}

	// Last page button
	tgButtons = append(tgButtons, tgbotapi.NewInlineKeyboardButtonData("Last page", c.createButtonInfo(dbDataSize-pageSize).String()))

	return tgButtons
}

func (c *BusinessWorkplaceCommander) createButtonInfo(dataOffset uint64) path.CallbackPath {
	serializedData, _ := json.Marshal(CallbackListData{
		Offset: dataOffset,
		Limit: 3,
	})

	var callbackPath = path.CallbackPath{
		Domain:       "business",
		Subdomain:    "workplace",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return callbackPath
}
