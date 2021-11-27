package ticket

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *TicketCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("Failed to parse data %v", callbackPath.CallbackData)

		return
	}

	outputMessage := ""

	tickets := c.ticketService.List(parsedData.Cursor, parsedData.Limit)

	for i, t := range tickets {
		outputMessage += fmt.Sprintf("%v. ", uint64(i)+parsedData.Cursor+1)
		outputMessage += fmt.Sprintf("User: %v,\nSchedule: %v", t.User, t.Schedule)
		outputMessage += "\n\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMessage)

	buttons := createButtonsIfNecessary(parsedData, len(tickets))
	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(buttons...),
		)
	}

	c.bot.Send(msg)
}

func createButtonsIfNecessary(parsedData CallbackListData, ticketsCount int) []tgbotapi.InlineKeyboardButton {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0, 2)
	if parsedData.Cursor > 0 {
		previousPageStart := parsedData.Cursor - ListLimit

		if previousPageStart < 0 {
			previousPageStart = 0
		}

		serializedData, _ := json.Marshal(
			CallbackListData{
				Cursor: previousPageStart,
				Limit:  ListLimit,
			})

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(PreviousPageText, CallbackListPrefix+string(serializedData)))
	}

	nextPageCursor := parsedData.Cursor + parsedData.Limit
	if ticketsCount == ListLimit {
		serializedData, _ := json.Marshal(
			CallbackListData{
				Cursor: nextPageCursor,
				Limit:  ListLimit,
			})

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(NextPageText, CallbackListPrefix+string(serializedData)))
	}

	return buttons
}
