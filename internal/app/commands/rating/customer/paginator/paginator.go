package paginator

import (
	"encoding/json"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/rating/customer"
)

const DefaultLimit = 2

// Paginator отвечает за вывод информации о покупателях порциями
type Paginator struct {
	customerService *customer.DummyService
	limit           uint64
}

type CallbackListData struct {
	Offset int `json:"offset"`
}

func NewPaginator(customerService *customer.DummyService, limit uint64) *Paginator {
	if limit == 0 {
		limit = DefaultLimit
	}
	return &Paginator{
		customerService: customerService,
		limit:           limit,
	}
}

func (p *Paginator) GetMessage(chatID int64, data CallbackListData) (msg tgbotapi.MessageConfig, err error) {
	outputMsgText := "Here all the customers: \n\n"
	customers, err := p.customerService.List(uint64(data.Offset), p.limit)
	if err != nil {
		return
	}

	for _, p := range customers {
		outputMsgText += fmt.Sprintf("%s\n", p)
	}

	var buttonRow []tgbotapi.InlineKeyboardButton

	buttonRow, err = p.appendPrevButton(buttonRow, data.Offset)
	if err != nil {
		return
	}
	buttonRow, err = p.appendNextButton(buttonRow, data.Offset)
	if err != nil {
		return
	}

	msg = tgbotapi.NewMessage(chatID, outputMsgText)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		buttonRow,
	)

	return
}

func (p *Paginator) appendNextButton(buttonRow []tgbotapi.InlineKeyboardButton, offset int) ([]tgbotapi.InlineKeyboardButton, error) {
	if offset+int(p.limit) < p.customerService.Count() {
		nextButton, err := getButton("Next page", offset+int(p.limit))
		if err != nil {
			return nil, err
		}
		buttonRow = append(buttonRow, nextButton)
	}

	return buttonRow, nil
}

func (p *Paginator) appendPrevButton(buttonRow []tgbotapi.InlineKeyboardButton, offset int) ([]tgbotapi.InlineKeyboardButton, error) {
	if offset > 0 {
		prevOffset := offset - int(p.limit)
		if prevOffset < 0 {
			prevOffset = 0
		}
		prevButton, err := getButton("Previous page", prevOffset)
		if err != nil {
			return nil, err
		}
		buttonRow = append(buttonRow, prevButton)
	}

	return buttonRow, nil
}

func getButton(title string, offset int) (button tgbotapi.InlineKeyboardButton, err error) {
	serializedData, err := json.Marshal(CallbackListData{
		Offset: offset,
	})
	if err != nil {
		return
	}

	callbackPath := path.CallbackPath{
		Domain:       "rating",
		Subdomain:    "customer",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return tgbotapi.NewInlineKeyboardButtonData(title, callbackPath.String()), nil
}
