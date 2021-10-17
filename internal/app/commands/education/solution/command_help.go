package solution

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *SolutionCommander) Help(inputMsg *tgbotapi.Message){
	c.SendMessage(inputMsg, OneCommandName("help") + " - Эта справка\n"+
		OneCommandName("get") + " - получить элемент по номеру\n"+
		OneCommandName("list") + " - получить список элементов\n"+
		OneCommandName("delete") + " - удалить элемент по номеру\n"+
		OneCommandName("new") + " - создать новый элемент\n"+
		OneCommandName("edit") + " - отредактировать элемент\n")
}
