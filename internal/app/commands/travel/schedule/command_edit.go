package schedule

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"unicode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/travel"
	"github.com/sirupsen/logrus"
)

func (c *TravelScheduleCommander) Edit(msg *tgbotapi.Message) {
	re := regexp.MustCompile(`^\s*\d+\s+.*$`)
	args := msg.CommandArguments()
	if !re.MatchString(args) {
		logrus.Errorf("wrong arguments: %s", msg.CommandArguments())
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Wrong arguments")
		c.bot.Send(reply)
		return
	}

	d := 0
	for unicode.IsSpace(rune(args[d])) {
		d++
	}

	for unicode.IsDigit(rune(args[d])) {
		d++
	}

	id, err := strconv.ParseUint(args[:d], 10, 64)
	if err != nil {
		logrus.Error(err)
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Not a valid id")
		c.bot.Send(reply)
		return
	}

	var schedule travel.Schedule
	if err := json.Unmarshal([]byte(args[d:]), &schedule); err != nil {
		logrus.Error(err)
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Wrong arguments")
		c.bot.Send(reply)
		return
	}

	if err := c.service.Update(id, schedule); err != nil {
		logrus.Error(err)
		reply := tgbotapi.NewMessage(msg.Chat.ID, "Could not edit entity")
		c.bot.Send(reply)
		return
	}

	text := fmt.Sprintf("Updated %d: %s", id, schedule.Name)
	reply := tgbotapi.NewMessage(msg.Chat.ID, text)
	c.bot.Send(reply)
}
