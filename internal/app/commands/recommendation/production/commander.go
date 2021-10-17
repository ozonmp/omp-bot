package production

import (
	"errors"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/recommendation/production"
)

const textWrong = "something went wrong"

const commanderName = "RecommendationProductionCommander"

var errWrongArgs = errors.New("wrong args")

type RecommendationProductionCommander struct {
	bot               *tgbotapi.BotAPI
	productionService ProductionService
}

func NewRecommendationProductionCommander(bot *tgbotapi.BotAPI) *RecommendationProductionCommander {
	service := production.NewRecommendationProductionService()

	return &RecommendationProductionCommander{
		bot:               bot,
		productionService: service,
	}
}

func (c *RecommendationProductionCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("%v.HandleCallback: unknown callback name: %v", commanderName, callbackPath.CallbackName)
	}
}

func (c *RecommendationProductionCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "create":
		c.Create(msg)
	case "read":
		c.Read(msg)
	case "update":
		c.Update(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}

func (c *RecommendationProductionCommander) sendMessage(msg tgbotapi.Chattable) (tgbotapi.Message, error) {
	retMsg, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("%v send message error %v", commanderName, err)
	}

	return retMsg, err
}
