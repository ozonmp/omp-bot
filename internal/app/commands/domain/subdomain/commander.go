package subdomain

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/domain/subdomain"
)

const (
	DomainName    = "domain"
	SubdomainName = "subdomain"
)

type SubdomainCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type DummySubdomainCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService service.SubdomainService
}

func NewSubdomainCommander(bot *tgbotapi.BotAPI, service service.SubdomainService) *DummySubdomainCommander {
	return &DummySubdomainCommander{
		bot:              bot,
		subdomainService: service,
	}
}

func (c *DummySubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DummySubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *DummySubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		log.Printf("DummySubdomainCommander.HandleCommand: unknown command name: %s", commandPath.CommandName)
	}
}
