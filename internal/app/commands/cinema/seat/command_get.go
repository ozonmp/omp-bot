package seat

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CinemaSeatCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	seat, err := c.subdomainService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		return
	}

	message := fmt.Sprintf("Seat #%d (Row: %d Number %d) price: %d", seat.ID, seat.Row, seat.Number, seat.Price)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		message,
	)

	c.bot.Send(msg)
}
