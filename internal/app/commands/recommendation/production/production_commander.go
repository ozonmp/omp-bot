package production

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type ProductionCommander interface {
	List(inputMsg *tgbotapi.Message)
	Create(inputMsg *tgbotapi.Message)
	Read(inputMsg *tgbotapi.Message)
	Update(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	Default(inputMsg *tgbotapi.Message)
}
