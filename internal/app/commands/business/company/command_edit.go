package company

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/business"
)

func (c *CompanyCommander) Edit(inputMessage *tgbotapi.Message) {
	idx := 0

	err := c.companyService.Update(uint64(idx), business.Company{})
	if err != nil {
		log.Printf("fail to edit company with idx %d: %v", idx, err)
	}
}
