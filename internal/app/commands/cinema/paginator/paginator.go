package paginator

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Paginator struct {
	Direction string `json:"direction"`
	Page      int    `json:"page"`
}

func NewPaginator(s string) *Paginator {
	callBackData := new(Paginator)
	_ = json.Unmarshal([]byte(s), callBackData)
	return callBackData
}

func (p Paginator) String() string {
	if str, err := json.Marshal(p); err == nil {
		return string(str)
	} else {
		return ""
	}
}

func (p Paginator) makeRow(name string) tgbotapi.InlineKeyboardButton {
	p.Direction = name
	return tgbotapi.NewInlineKeyboardButtonData(name, "cinema__film__list__"+p.String())
}

func (p Paginator) NewKeyBoard() *tgbotapi.InlineKeyboardMarkup {
	keyBoard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			p.makeRow("next")),
		tgbotapi.NewInlineKeyboardRow(
			p.makeRow("prev")),
	)
	return &keyBoard
}
