package championat

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *ChampionatCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Wrong ID: it should be a positive integer number",
		)
		_, err = c.bot.Send(msg)
		return
	}
	if idx < 0 {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Wrong ID: it should be a positive number",
		)
		_, err = c.bot.Send(msg)
		return
	}

	err = c.championatService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to delete championat with id %v: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Championat with id "+strconv.Itoa(idx)+" was deleted!",
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ChampionatCommander.Describe: error sending reply message to chat - %v", err)
	}
}
