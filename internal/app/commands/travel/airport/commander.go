package airport

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/travel/airport"
	"log"
)

type TravelAirportCommander struct {
	bot            *tgbotapi.BotAPI
	airportService *airport.Service
}

func NewTravelAirportCommander(
	bot *tgbotapi.BotAPI) *TravelAirportCommander {
	return &TravelAirportCommander{bot, airport.NewService()}
}

func (c *TravelAirportCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	default:
		log.Printf("TravelAirportCommander got %s (%s)}",
			callbackPath.CallbackName,
			callback,
		)
	}
}

func (c *TravelAirportCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	default:
		c.Default(msg)
	}
}
