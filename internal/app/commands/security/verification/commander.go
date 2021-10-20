package verification

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/security/verification"
)

type SecurityVerificationCommander struct {
	bot                 *tgbotapi.BotAPI
	verificationService VerificationService
}

type VerificationService interface {
	Describe(VerificationID uint64) (*verification.Verification, error)
	List(cursor uint64, limit uint64) ([]verification.Verification, error)
	Create(verification.Verification) (uint64, error)
	Update(verificationID uint64, verification verification.Verification) error
	Remove(verificationID uint64) (bool, error)
}

func NewSecurityVerificationCommander(bot *tgbotapi.BotAPI) *SecurityVerificationCommander {
	verificationService := verification.NewVerificationService([]verification.Verification{
		{Title: "one"},
		{Title: "two"},
		{Title: "three"},
		{Title: "four"},
		{Title: "five"},
	})

	return &SecurityVerificationCommander{
		bot:                 bot,
		verificationService: verificationService,
	}
}

func (c *SecurityVerificationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	default:
		log.Printf("SecurityVerificationCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *SecurityVerificationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
	}
}

var internalError = "internal error"

func (c *SecurityVerificationCommander) sendErrorMsg(commandName string, msg tgbotapi.MessageConfig) {
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("SecurityVerificationCommander.%s: error sending error message to chat - %v", commandName, err)
	}
}
