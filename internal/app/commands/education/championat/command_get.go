package championat

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ChampionatCommander) Get(inputMessage *tgbotapi.Message) {
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

	championat, err := c.championatService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get championat with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		championat.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ChampionatCommander.Describe: error sending reply message to chat - %v", err)
	}
}
