package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/servicedata"
	"log"
)

func (c *SolutionCommander) Default(inputMessage *tgbotapi.Message) {
	if idx, ok := servicedata.EditedChat[inputMessage.Chat.ID]; ok {
		switch idx.OperationType {
		case servicedata.NewoperationData:
			c.newCommit(inputMessage)
		case servicedata.EditoperationData:
			c.editCommit(inputMessage)
		default:
			c.SendMessage(inputMessage, "Неизвстная комманда " + idx.String())
		}
	} else {
		log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
		c.SendMessage(inputMessage, "You wrote: "+inputMessage.Text)
	}
}
