package recommendation

import (
	service "github.com/ozonmp/omp-bot/internal/service/recommendation/product"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	commander "github.com/ozonmp/omp-bot/internal/app/commands/recommendation/product"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type RecommendationCommander struct {
	bot              *tgbotapi.BotAPI
	productCommander commander.Commander
}

func NewRecommendationCommander(
	bot *tgbotapi.BotAPI,
) *RecommendationCommander {
	service := service.NewDummyProductService()
	return &RecommendationCommander{
		bot:              bot,
		productCommander: commander.NewProductCommander(bot, service),
	}
}

func (commander *RecommendationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case Product:
		commander.productCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (commander *RecommendationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case Product:
		commander.productCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
