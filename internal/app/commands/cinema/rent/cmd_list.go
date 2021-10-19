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
	outputMsgText := "Films and serials prices:\n\n"

	items, err := c.service.List(cursor, limit)
	if err != nil {
		log.Printf("CinemaRentCommander.List: %v", err)
		msg := tgbotapi.NewMessage(ChatID, fmt.Sprintf("%v", err))
		if _, err := c.bot.Send(msg); err != nil {
			log.Printf("CinemaRentCommander.List: service list error inform: %v", err)
		}
		return
	}

	for _, item := range items {
		outputMsgText += fmt.Sprintf("%d => %s\n", item.RecordIndex, item.String())
	}

	msg := tgbotapi.NewMessage(ChatID, outputMsgText)

	buttons := []tgbotapi.InlineKeyboardButton{}

	if cursor != 0 {
		buttons = append(buttons, c.listBackButton(cursor, limit))
	}

	// Если количество отобранных элементов равно лимиту, проверим
	// есть ли элемнты за лимитом, и если есть, отобразим кнопку следующей страницы
	if uint64(len(items)) == limit {
		nextItems, _ := c.service.List(cursor+limit+1, limit)
		if len(nextItems) > 0 {
			buttons = append(buttons, c.listNextButton(cursor, limit))
		}
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
	cursor = cursor - limit - 1
	if cursor < 0 {
		// защита от "дурака"
		cursor = 0
	}

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
		Cursor: int64(cursor + limit + 1),
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
