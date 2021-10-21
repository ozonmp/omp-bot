package rent
// TODO
//import (
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
//	"log"
//)
//
//var command_default = func() CommandMetadata {
//	deleteMetadata := newCommandMetadata("default", "just talks")
//	deleteMetadata.Exec = func (c *Commander, chatId int64, command string, commandArgs string) {
//		//log.Printf("[%s] %s", message.From.UserName, message.Text)
//
//		msg := tgbotapi.NewMessage(chatId, "You wrote:"+message.Text)
//		msg.ReplyToMessageID = message.MessageID
//
//		var args string
//		appendNewOrEditButtons(&msg, args)
//
//		c.sendWithLog(msg, "DefaultCommand")
//	}
//}()