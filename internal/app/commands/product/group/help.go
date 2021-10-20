package group

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *ProductGroupCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		`
		/help__product__group - Print help
		/get__product__group 'id' - get group by 'id'
		/list__product__group - groups list (3 per page) 
		/delete__product_group 'id' - delete group by 'id'
		/new__product__group 'owner' 'item' - create group by 'Owner' and 'Item'
		/edit__product__group 'id' 'owner' ítem' - edit group set 'Owner' 'Item' in Group 'id'
		`,
	)

	c.bot.Send(msg)
}
