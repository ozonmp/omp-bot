package verification

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/security/verification"
)

func (c *SecurityVerificationCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		c.sendErrorMsg("Get", tgbotapi.NewMessage(inputMessage.Chat.ID, "wrong args expected one number, got:"+args))
		return
	}

	product, err := c.verificationService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		if err == verification.ErrEntityNotExists {
			c.sendErrorMsg("Get", tgbotapi.NewMessage(inputMessage.Chat.ID, "Product with id:"+args+" does not exist"))
			return
		}
		c.sendErrorMsg("Get", tgbotapi.NewMessage(inputMessage.Chat.ID, internalError))
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("SecurityVerificationCommander.Get: error sending reply message to chat - %v", err)
	}
}
