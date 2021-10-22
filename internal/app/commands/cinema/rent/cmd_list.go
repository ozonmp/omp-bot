package rent

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *CinemaRentCommander) List(inputMessage *tgbotapi.Message) {
	c.sendList(inputMessage.Chat.ID, 0, 10)
}

func (c *CinemaRentCommander) sendList(ChatID int64, cursor, limit uint64) {
	log.Printf("CinemaRentCommander.sendList: cursor = %v, count = %v", cursor, limit)
	outputMsgText := "Films and serials prices:\n\n"

	// Запроосим количество элементов сверх лимита (чтоб понять есть ли элементы дальше)
	items, err := c.service.List(cursor, limit+1)
	if err != nil {
		log.Printf("CinemaRentCommander.List: %v", err)
		msg := tgbotapi.NewMessage(ChatID, fmt.Sprintf("%v", err))
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("CinemaRentCommander.List: service list error inform: %v", err)
		}
		return
	}

	for j, item := range items {
		if uint64(j) == limit {
			break
		}
		outputMsgText += fmt.Sprintf("%d => %s\n", item.RecordIndex, item.String())
	}

	msg := tgbotapi.NewMessage(ChatID, outputMsgText)

	buttons := []tgbotapi.InlineKeyboardButton{}

	if cursor != 0 {
		nextCursor := cursor - limit
		if nextCursor == 1 {
			nextCursor = 0
		}
		buttons = append(buttons, c.listBackButton(nextCursor, limit))
	}

	// Если количество отобранных элементов больше лимита, отобразим кнопку следующей страницы
	if uint64(len(items)) > limit {
		nextCursor := cursor + limit
		if cursor == 0 {
			nextCursor += 1
		}
		buttons = append(buttons, c.listNextButton(nextCursor, limit))
	}

	if len(buttons) > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				buttons...,
			),
		)
	}

	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("CinemaRentCommander.List: %v", err)
	}

}

func (c *CinemaRentCommander) listBackButton(cursor, limit uint64) tgbotapi.InlineKeyboardButton {
	serializedData, _ := json.Marshal(ListLimitDTO{
		Cursor: int64(cursor),
		Limit:  int64(limit),
	})

	callbackPath := path.CallbackPath{
		Domain:       "cinema",
		Subdomain:    "rent",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPath.String())
}

func (c CinemaRentCommander) listNextButton(cursor, limit uint64) tgbotapi.InlineKeyboardButton {
	serializedData, _ := json.Marshal(ListLimitDTO{
		Cursor: int64(cursor),
		Limit:  int64(limit),
	})

	callbackPath := path.CallbackPath{
		Domain:       "cinema",
		Subdomain:    "rent",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String())
}
