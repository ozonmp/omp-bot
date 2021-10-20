package verification

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/security/verification"
)

const MaxUint = ^uint64(0)

func (c *SecurityVerificationCommander) List(inputMessage *tgbotapi.Message) {
	//list should return all values
	products, err := c.verificationService.List(0, MaxUint)
	if err != nil {
		log.Printf("fail to get list of products: %v", err)
		c.sendErrorMsg("List", tgbotapi.NewMessage(inputMessage.Chat.ID, internalError))
		return
	}

	outputMsg, err := c.createListOutputMessage(products)
	if err != nil {
		log.Printf("fail to create string from list of products: %v", err)
		c.sendErrorMsg("List", tgbotapi.NewMessage(inputMessage.Chat.ID, internalError))
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.List: error sending reply message to chat - %v", err)
	}
}

func (c SecurityVerificationCommander) createListOutputMessage(products []verification.Verification) (string, error) {
	b := strings.Builder{}
	if _, err := b.WriteString("Here all the products: \n\n"); err != nil {
		return "", err
	}
	for _, p := range products {
		if _, err := b.WriteString(p.String() + "\n"); err != nil {
			return "", err
		}
	}
	return b.String(), nil
}
