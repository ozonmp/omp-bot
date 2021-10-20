package business

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/ozonmp/omp-bot/internal/app/commands/business/company"
	service "github.com/ozonmp/omp-bot/internal/service/business/company"
)

type CompanyCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

func NewCompanyCommander(bot *tgbotapi.BotAPI, service service.CompanyService) CompanyCommander {
	return company.NewCompanyCommander(bot)
}
