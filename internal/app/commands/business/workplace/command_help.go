package workplace

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BusinessWorkplaceCommander) Help(inputMessage *tgbotapi.Message) {
	var msg = tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__business__workplace — print list of commands\n"+
			"/get__business__workplace — get a entity\n"+
			"/list__business__workplace — get a list of your entity\n"+
			"/delete__business__workplace — delete an existing entity\n"+
			"/new__business__workplace — create a new entity\n"+
			"/edit__business__workplace — edit a entity\n",
	)

	c.bot.Send(msg)
}
