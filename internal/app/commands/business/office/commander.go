package office

import (
	"github.com/ozonmp/omp-bot/internal/model/business"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/business/office"
)

type OfficeService interface {
	Describe(officeId uint64) (*business.Office, error)
	List(cursor uint64, limit uint64) ([]business.Office, error)
	Create(o business.Office) (uint64, error)
	Update(officeId uint64, office business.Office) error
	Remove(officeId uint64) (bool, error)
}

type OfficeCommander struct {
	bot           *tgbotapi.BotAPI
	officeService OfficeService
}

func NewOfficeCommander(
	bot *tgbotapi.BotAPI,
) *OfficeCommander {
	officeService := office.NewDummyOfficeService()

	return &OfficeCommander{
		bot:           bot,
		officeService: officeService,
	}
}

func (c *OfficeCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("OfficeCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *OfficeCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "create":
		c.Create(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Help(msg)
	}
}
