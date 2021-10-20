package certificate

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/loyalty/certificate"
	"log"
	"strconv"
)

func (c *LoyaltyCertificateCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	outputMsg := ""
	parsedData := CertificateData{}
	err := json.Unmarshal([]byte(args), &parsedData)
	if err != nil {
		outputMsg = "Pass valid JSON serialized data as parameter"
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			outputMsg,
		)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("LoyaltyCertificateCommander.Edit: error sending reply message to chat - %v", err)
		}
		return
	}

	newCertificate := certificate.Certificate{
		Id:          parsedData.Id,
		SellerTitle: parsedData.SellerTitle,
		Amount:      parsedData.Amount,
		ExpireDate:  parsedData.ExpireDate,
	}

	err = c.certificateService.Update(newCertificate.Id, newCertificate)
	outputMsg = "Certificate with ID " + strconv.Itoa(int(newCertificate.Id))
	if err != nil {
		log.Printf("failed to update certificate with id %d: %v", newCertificate.Id, err)
		outputMsg += " was not found"
	} else {
		outputMsg += " was updated"
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsg,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.Edit: error sending reply message to chat - %v", err)
	}
}
