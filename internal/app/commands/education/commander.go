package education

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/education/solution"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service_consts"
	"log"
)

type EducationCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type Education_Commander struct {
	bot                *tgbotapi.BotAPI
	SolutionCommander solution.Solution_Commander
}

func NewEducationCommander(
	bot *tgbotapi.BotAPI,
) *Education_Commander {
	return &Education_Commander{
		bot: bot,
		// SolutionCommander
		SolutionCommander: solution.NewSolutionCommander(bot),
	}
}

func (c *Education_Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case service_consts.Solution:
		c.SolutionCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("Education_Commander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *Education_Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case service_consts.Solution:
		c.SolutionCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("Education_Commander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
