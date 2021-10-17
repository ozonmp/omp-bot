package solution

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *SolutionCommander) Get(inputMsg *tgbotapi.Message){
	TextMsg := ""
	defer func() {
		c.SendMessage(inputMsg, TextMsg)
	}()
	idx, TextMsg := GetArgument(inputMsg)
	if TextMsg != "" { return}

	product, err := c.SolutionService.Describe(idx)
	if err != nil {
		TextMsg = fmt.Sprintf("fail to get product with idx %d: %v", idx, err)
		log.Println(TextMsg)
		return
	}
	TextMsg = product.String()
}

