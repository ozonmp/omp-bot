package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/servicedata"
)

func (c *SolutionCommander) New(inputMsg *tgbotapi.Message){
	servicedata.EditedChat[inputMsg.Chat.ID] = *(servicedata.GetOperationData(0, servicedata.NewoperationData))
	TextMsg := "Новая запись должна содержать поля TaskID, Autor, Title. Все поля "+
		"должны быть в одном сообщении каждое поле в отдельной строке."
	c.SendMessage(inputMsg, TextMsg)
}

