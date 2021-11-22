package ticket

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Sender interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}
