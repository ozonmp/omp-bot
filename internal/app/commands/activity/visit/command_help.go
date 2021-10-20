package visit

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *VisitCommanderStruct) Help(inputMessage *tgbotapi.Message) {
	c.Send(
		inputMessage.Chat.ID,
		"/help__activity__visit - visit help\n"+
			"/list__activity__visit - list visits\n"+
			"/get__activity__visit - get visit\n"+
			"/new__activity__visit - new visit\n"+
			"/edit__activity__visit - edit visit\n"+
			"/delete__activity__visit - delete visit",
	)
}
