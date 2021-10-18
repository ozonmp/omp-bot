package film

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaFilmCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__cinema__film - help\n"+
			"/list__cinema__film - list films\n"+
			"/get__cinema__film - get film\n"+
			"/delete__cinema__film - get film\n"+
			"/new__cinema__film - create new film\n"+
			"/edit__cinema__film - edit film\n",
	)

	_ = c.sendMessage(msg)
}
