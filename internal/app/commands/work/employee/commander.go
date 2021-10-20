package employee

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/work/employee"
)

type DemoSubdomainCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService employee.EmployeeService
}

func NewDemoSubdomainCommander(
	bot *tgbotapi.BotAPI,
) *DemoSubdomainCommander {
	service := employee.NewService()

	return &DemoSubdomainCommander{
		bot:              bot,
		subdomainService: service,
	}
}

func (c *DemoSubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *DemoSubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Update(msg)
	default:
		c.Default(msg)
	}
}

func (c *DemoSubdomainCommander) sendMessage(input *tgbotapi.Message, text string) {
	_, err := c.bot.Send(tgbotapi.NewMessage(input.Chat.ID, text))
	if err != nil {
		log.Printf("Error sending reply message to chat - %v", err)
	}
}

func (c *DemoSubdomainCommander) sendWrongArgs(input *tgbotapi.Message) {
	c.sendMessage(input, "Error -> Wrong Args")
}
