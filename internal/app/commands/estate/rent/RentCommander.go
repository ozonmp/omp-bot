package rent

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type RentCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return sendError not implemented
	Edit(inputMsg *tgbotapi.Message) // return sendError not implemented
}
