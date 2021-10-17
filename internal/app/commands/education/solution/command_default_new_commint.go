package solution

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"github.com/ozonmp/omp-bot/internal/servicedata"
	"strconv"
	"strings"
)

func (c *SolutionCommander) newCommit(inputMessage *tgbotapi.Message) {
	TextMsg := ""
	defer func() {
		c.SendMessage(inputMessage, TextMsg)
	}()
	data := strings.Split(inputMessage.Text, "\n")
	if len(data) != 3 {
		TextMsg = "В сообщение должно быть 3 строки, повторите ввод, пожалуйста"
		return
	}
	taskID, err := strconv.ParseUint(data[0], 0, 64)
	if err != nil {
		TextMsg = "Первая строка не содержит число, повторите ввод, пожалуйста"
		return
	}
	idx, _ := servicedata.EditedChat[inputMessage.Chat.ID]
	NewID := c.SolutionService.CreateNewID()
	idx.ProductID = NewID
	solution := education.Solution{	}
	solution.Id = NewID
	solution.TaskID = taskID
	solution.Autor = data[1]
	solution.Title = data[2]
	c.SolutionService.Create(NewID, solution)
	delete(servicedata.EditedChat, inputMessage.Chat.ID)
	sol, _ := c.SolutionService.Describe(idx.ProductID)
	TextMsg = "Запись добавлена: \n " + sol.String()
}

