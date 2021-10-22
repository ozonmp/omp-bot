package rent

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/estate/rent"
	"log"
	"strings"
)

func NewRentCommander(bot *tgbotapi.BotAPI, service service.RentService) *Commander {
	commander := &Commander{bot: bot, productService: service, offset: 5}
	return commander
}

type Commander struct {
	domain         string
	subdomain      string
	bot            *tgbotapi.BotAPI
	productService service.RentService
	offset         uint64
}

func (c *Commander) IsBelongsSubdomain(subdomain string) bool {
	return strings.HasPrefix(subdomain, c.subdomain)
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	command := CommandName(callbackPath.CallbackName)
	switch command {
	case CmndList:
		c.CallbackList(callback, callbackPath)
	case CmndEdit:
		c.CallbackEdit(callback, callbackPath)
	case CmndNew:
		c.CallbackNew(callback, callbackPath)
	case CmndDelete:
		c.CallbackDelete(callback, callbackPath)
	default:
		log.Printf("EstateRestCommander.HandleCallback: unknown callback Name: %s", callbackPath.CallbackName)
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	command := CommandName(commandPath.CommandName)
	switch command {
	case CmndHelp:
		c.Help(msg)
	case CmndList:
		c.List(msg)
	case CmndGet:
		c.Get(msg)
	case CmndEdit:
		c.Edit(msg)
	case CmndNew:
		c.New(msg)
	case CmndDelete:
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}

func (c *Commander) Delete(inputMsg *tgbotapi.Message) {
	command_delete.Exec(c, inputMsg.Chat.ID, inputMsg.Command(), inputMsg.CommandArguments())
}

func (c *Commander) Edit(inputMsg *tgbotapi.Message) {
	command_edit.Exec(c, inputMsg.Chat.ID, inputMsg.Command(), inputMsg.CommandArguments())
}

func (c *Commander) Get(inputMsg *tgbotapi.Message) {
	command_get.Exec(c, inputMsg.Chat.ID, inputMsg.Command(), inputMsg.CommandArguments())
}

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	command_help.Exec(c, inputMsg.Chat.ID, inputMsg.Command(), inputMsg.CommandArguments())
}

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	command_list.Exec(c, inputMsg.Chat.ID, inputMsg.Command(), inputMsg.CommandArguments())
}

func (c *Commander) New(inputMsg *tgbotapi.Message) {
	command_new.Exec(c, inputMsg.Chat.ID, inputMsg.Command(), inputMsg.CommandArguments())
}

func (c *Commander) Default(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	command, commandArgs := parseCommandFromText(message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote:"+message.Text)
	msg.ReplyToMessageID = message.MessageID

	var newArgs string
	var editArgs string
	switch command {
	case CmndNew:
		newArgs = fmt.Sprintf(`%s__%s`, CmndNew.ToCallbackPath(), commandArgs)
		appendNewOrEditButtons(&msg, newArgs, editArgs)
	case CmndEdit:
		editArgs = fmt.Sprintf(`%s__%s`, CmndEdit.ToCallbackPath(), commandArgs)
		appendNewOrEditButtons(&msg, newArgs, editArgs)
	default:

	}

	c.sendWithLog(msg, "DefaultCommand")
}

func parseCommandFromText(msgText string) (CommandName, string) {
	if strings.HasPrefix(msgText, "@") {
		newRowIndex := strings.Index(msgText, "\n")
		msgText = msgText[newRowIndex+1:]

		if strings.HasPrefix(msgText, "/") {
			groundIndex := strings.Index(msgText, "_")
			spaceIndex := strings.Index(msgText, " ")
			return CommandName(msgText[1 : groundIndex-1]), msgText[spaceIndex+1:]
		}
	}

	return "", ""
}

func (c *Commander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	chatId := callback.Message.Chat.ID
	callerName := "CallbackList"

	args := commandListArgs{}
	if err := args.Unmarshal(callbackPath.CallbackData); err == nil {
		command_list.Exec(c, chatId, "", callbackPath.CallbackData)
	} else {
		c.sendError(chatId, err, callerName)
	}
}

func (c *Commander) CallbackEdit(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	chatId := callback.Message.Chat.ID
	command_edit.Exec(c, chatId, command_edit.Name.ToDomainCommand(), callbackPath.CallbackData)
}

func (c *Commander) CallbackNew(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	chatId := callback.Message.Chat.ID
	command_new.Exec(c, chatId, command_new.Name.ToDomainCommand(), callbackPath.CallbackData)
}

func (c *Commander) CallbackDelete(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	chatId := callback.Message.Chat.ID

	args := commandDeleteArgs{}
	if err := args.Unmarshal(callbackPath.CallbackData); err == nil {
		ok, deleteErr := c.productService.Remove(args.ID)
		if deleteErr != nil {
			c.sendError(chatId, deleteErr, "CallbackDelete")
			return
		}

		var msgTxt string
		if ok {
			msgTxt = fmt.Sprintf(`❌ rent deleted (ID: %v)
show %s`, args.ID, CmndList.ToDomainCommand())
		} else {
			msgTxt = fmt.Sprintf(`❗ rent not found (ID: %v)
show %s`, args.ID, CmndList.ToDomainCommand())
		}

		msg := tgbotapi.NewMessage(chatId, msgTxt)
		c.sendWithLog(msg, "CallbackDelete")
	}
}

func (c *Commander) sendError(chatID int64, err error, callerName string) bool {
	errMsg := fmt.Sprintf(`%s 
try %s`, err.Error(), CmndHelp.ToDomainCommand())
	msg := tgbotapi.NewMessage(chatID, errMsg)
	return c.sendWithLog(msg, callerName)
}

func (c *Commander) sendWithLog(msg tgbotapi.MessageConfig, callerName string) bool {
	if _, err := c.bot.Send(msg); err != nil {
		log.Printf("%s: error sending message %v", callerName, err)
		return false
	}

	return true
}
