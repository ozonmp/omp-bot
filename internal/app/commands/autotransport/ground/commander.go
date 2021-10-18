package ground

import (
	"errors"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/autotransport/ground"
)

type GroundCommander struct {
	bot           *tgbotapi.BotAPI
	groundService *ground.DummyGroundService
}

func NewGroundCommander(bot *tgbotapi.BotAPI) *GroundCommander {
	groundService := ground.NewDummyGroundService()

	return &GroundCommander{
		bot:           bot,
		groundService: groundService,
	}
}

func (c *GroundCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("GroundCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *GroundCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	var err error
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		err = c.List(msg)
	case "get":
		err = c.Get(msg)
	case "delete":
		err = c.Delete(msg)
	case "new":
		err = errors.New("not yet implemented")
	case "edit":
		err = errors.New("not yet implemented")
	default:
		c.Default(msg)
	}
	if err == nil {
		return
	}
	var userError error
	var outMsg tgbotapi.MessageConfig
	if errors.As(err, &userError) {
		outMsg = tgbotapi.NewMessage(msg.Chat.ID, "Input data were wrong: "+userError.Error())
	} else {
		log.Printf("Internal error %v", err)
		outMsg = tgbotapi.NewMessage(msg.Chat.ID, "Got internal exception")
	}

	if _, err := c.bot.Send(outMsg); err != nil {
		log.Printf("Send message errro %v", err)
	}

}
