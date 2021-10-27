package author

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type LicenseAuthorService interface {
	List(cursor uint64, limit uint64) ([]Author, error)
	Count() int
	Describe(idx uint64) (*Author, error)
	Remove(idx uint64) (bool, error)
	Create(author Author) (uint64, error)
	Update(idx uint64, author Author) error
}

type LicenseAuthorCommander struct {
	bot           *tgbotapi.BotAPI
	authorService *DummyAuthorService
}

func NewLicenseAuthorCommander(
	bot *tgbotapi.BotAPI,
) *LicenseAuthorCommander {
	authorService := NewDummyAuthorService()

	return &LicenseAuthorCommander{
		bot:           bot,
		authorService: authorService,
	}
}

func (c *LicenseAuthorCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("LicenseAuthorCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LicenseAuthorCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(0, msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		log.Printf("Invalid command %s", commandPath.CommandName)
	}
}
