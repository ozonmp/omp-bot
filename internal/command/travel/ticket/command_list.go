package ticket

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TicketCommander) List(inputMessage *tgbotapi.Message) {
	outputMessage := "All the tickets: \n\n"

	tickets, total := c.ticketService.List(0, ListLimit)

	for i, t := range tickets {
		outputMessage += fmt.Sprintf("%v. ", i+1)
		outputMessage += fmt.Sprintf("ID: %v,\nUser: %v,\nSchedule: %v", t.ID, t.User, t.Schedule)
		outputMessage += "\n\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)

	if total > ListLimit {
		serializedData, _ := json.Marshal(
			CallbackListData{
				Cursor: ListLimit,
				Limit:  ListLimit,
			})

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(NextPageText, CallbackListPrefix+string(serializedData)),
			),
		)
	}

	c.bot.Send(msg)
}
