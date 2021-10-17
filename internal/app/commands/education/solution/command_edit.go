package solution

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/servicedata"
	"log"
)

func (c *SolutionCommander) Edit(inputMsg *tgbotapi.Message){
	TextMsg := ""
	defer func() {
		c.SendMessage(inputMsg, TextMsg)
	}()
	idx, TextMsg := GetArgument(inputMsg)
	if TextMsg != "" { return}

	product, err := c.SolutionService.Describe(idx)
	if err != nil {
		TextMsg = fmt.Sprintf("fail to get product with idx %d: %v", idx, err)
		log.Println(TextMsg)
		return
	}
	servicedata.EditedChat[inputMsg.Chat.ID] = *(servicedata.GetOperationData(idx, servicedata.EditoperationData))
	TextMsg = product.String() + "\n Измененная запись должна содержать поля TaskID, Autor, Title. Все поля "+
		"должны быть в одном сообщении каждое поле в отдельной строке."
}

