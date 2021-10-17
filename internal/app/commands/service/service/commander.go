package service

import (
	"errors"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/service/service"
)

type ServiceServiceCommander struct {
	bot            *tgbotapi.BotAPI
	serviceService *service.Service
}

func NewServiceServiceCommander(
	bot *tgbotapi.BotAPI,
) *ServiceServiceCommander {
	serviceService := service.NewService()

	return &ServiceServiceCommander{
		bot:            bot,
		serviceService: serviceService,
	}
}

func (c *ServiceServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ServiceServiceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	var err error

	log.Printf("[%s] %s", msg.From.UserName, msg.Text)

	switch commandPath.CommandName {
	case "help":
		_, err = c.Help(msg)
	case "list":
		_, err = c.List(msg)
	case "get":
		_, err = c.Get(msg)
	case "new":
		_, err = c.Create(msg)
	case "edit":
		_, err = c.Edit(msg)
	case "delete":
		_, err = c.Delete(msg)
	default:
		_, err = c.Default(msg)
	}

	if err != nil {
		c.HandleError(msg, err)
	}
}

var badRequest *service.BadRequestError

func (c *ServiceServiceCommander) HandleError(msg *tgbotapi.Message, err error) {
	if err == nil {
		return
	}

	var errMessage string
	if errors.As(err, &badRequest) {
		errMessage = fmt.Sprintf("Service.Service BadRequestError: %v", err)
	} else {
		errMessage = fmt.Sprintf("Service.Service BadRequestError: %v", err)
	}

	log.Printf(errMessage)

	outgoingMessage := tgbotapi.NewMessage(msg.Chat.ID, errMessage)
	_, sendingErr := c.bot.Send(outgoingMessage)

	if sendingErr != nil {
		log.Printf("Service.Service unhandled error: ", sendingErr)
	}
}
