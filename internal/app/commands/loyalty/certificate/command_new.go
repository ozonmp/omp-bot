package certificate

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/loyalty/certificate"
	"log"
	"strconv"
	"time"
)

func (c *LoyaltyCertificateCommander) New(inputMessage *tgbotapi.Message) {

	lastIndex := c.certificateService.Certificates[len(c.certificateService.Certificates) - 1].Id

	newCertificate := certificate.Certificate{
		Id:          lastIndex + 1,
		SellerTitle: "New Seller",
		Amount:      5000,
		ExpireDate:  time.Now().AddDate(0, 1, 0),
	}

	newId, err := c.certificateService.Create(newCertificate)
	outputMsg := "Certificate with ID " + strconv.Itoa(int(newId))
	if err != nil {
		log.Printf("failed to create certificate with id %d: %v", newCertificate.Id, err)
		outputMsg += " already exists"
	} else {
		outputMsg += " was created"
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsg,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.Get: error sending reply message to chat - %v", err)
	}
}
