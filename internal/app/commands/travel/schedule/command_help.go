package schedule

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const helpText = `/get__travel__schedule <id> - get entity by id
/list__travel__schedule - list entities paginated by 10
/delete__travel__schedule <id> - delete entity by id
/new__travel__schedule <entity> - create entity using json. Exapmle entity: {"name": "example"}
/edit__travel__schedule <id> <entity> - edit entity. Exapmle entity: {"name": "edit"}
/help__travel__schedule - this help`

func (c *TravelScheduleCommander) Help(msg *tgbotapi.Message) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, helpText)
	c.bot.Send(reply)
}
