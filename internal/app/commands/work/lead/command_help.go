package lead

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LeadCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__work__lead - help\n"+
			"/get__work__lead - get a lead\n"+
			"/list__work__lead - all leads\n"+
			"/delete__work__lead - delete a lead\n"+
			"/new__work__lead - create a lead\n"+
			"/edit__work__lead - edit a lead",
	)

	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/help__work__lead"),
			tgbotapi.NewKeyboardButton("/list__work__lead"),
			tgbotapi.NewKeyboardButton("/new__work__lead"),
			tgbotapi.NewKeyboardButton("/edit__work__lead"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/get__work__lead"),
			tgbotapi.NewKeyboardButton("/get__work__lead 3"),
			tgbotapi.NewKeyboardButton("/get__work__lead 10000"),
			tgbotapi.NewKeyboardButton("/delete__work__lead"),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Work.LeadCommander.Help: error sending reply message to chat - %v", err)
	}
}
