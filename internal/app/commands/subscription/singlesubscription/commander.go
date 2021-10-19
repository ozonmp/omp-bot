package singlesubscription

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/subscription/singlesubscription"
)

const (
	DomainName    = "subscription"
	SubdomainName = "singleSubscription"
)

type SingleSubscriptionCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type DummySingleSubscriptionCommander struct {
	bot     *tgbotapi.BotAPI
	service service.SingleSubscriptionService
}

func NewSingleSubscriptionCommander(bot *tgbotapi.BotAPI, service service.SingleSubscriptionService) *DummySingleSubscriptionCommander {
	return &DummySingleSubscriptionCommander{
		bot:     bot,
		service: service,
	}
}

func (c *DummySingleSubscriptionCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DummySingleSubscriptionCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *DummySingleSubscriptionCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
		log.Printf("DummySingleSubscriptionCommander.HandleCommand: unknown command name: %s", commandPath.CommandName)
	}
}
