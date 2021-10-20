package card

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (p *DummyCardCommander) Help(inMsg *tgbotapi.Message){
	msg := tgbotapi.NewMessage(inMsg.Chat.ID,
	"/help__bank__card - list of commands\n" +
	"/get__bank__card - get an existing card\n" +
		"/list__bank__card - get a list of cards\n" +
		"/delete__bank__card - delete an existing card\n",
	)

	p.bot.Send(msg)
}
