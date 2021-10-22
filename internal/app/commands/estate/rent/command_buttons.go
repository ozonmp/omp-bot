package rent

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/estate"
)

func templateFactory(
	command string,
	rent *estate.Rent,
) *string {
	const template = "fill:\n%s\n%s"
	res := fmt.Sprintf(
		template,
		command,
		rent.ToShortRowsString(),
	)
	return &res
}

func appendEditAndDeleteButtons(msg *tgbotapi.MessageConfig, editArgs string, deleteArgs string) {
	btnEdit := tgbotapi.NewInlineKeyboardButtonData("✏ edit", "")
	btnEdit.SwitchInlineQueryCurrentChat = &editArgs

	btnDelete := tgbotapi.NewInlineKeyboardButtonData("🗑 delete", deleteArgs)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(btnEdit, btnDelete),
	)
}

func appendPaginationButtons(msg *tgbotapi.MessageConfig, prevArgs string, nextArgs string) {
	getPaginationButtons := func(prevArgs string, nextArgs string) []tgbotapi.InlineKeyboardButton {
		buttons := make([]tgbotapi.InlineKeyboardButton, 0, 2)

		if len(prevArgs) > 0 {
			btnPrev := tgbotapi.NewInlineKeyboardButtonData("Prev page", prevArgs)
			buttons = append(buttons, btnPrev)
		}

		if len(nextArgs) > 0 {
			btnNext := tgbotapi.NewInlineKeyboardButtonData("Next page", nextArgs)
			buttons = append(buttons, btnNext)
		}

		return buttons
	}

	paginationButtons := getPaginationButtons(prevArgs, nextArgs)

	if len(paginationButtons) == 0 {
		return
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(paginationButtons...),
	)
}

func appendCreateNewButtons(msg *tgbotapi.MessageConfig, command string) {
	btnCar := tgbotapi.NewInlineKeyboardButtonData(estate.Car.Full(), "")
	btnHouse := tgbotapi.NewInlineKeyboardButtonData(estate.House.Full(), "")

	btnCar.SwitchInlineQueryCurrentChat = templateFactory(
		command,
		estate.EmptyRent(estate.Car),
	)
	btnHouse.SwitchInlineQueryCurrentChat = templateFactory(
		command,
		estate.EmptyRent(estate.House),
	)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(btnHouse, btnCar),
	)
}

func appendNewOrEditButtons(msg *tgbotapi.MessageConfig, newArgs string, editArgs string) {
	btnNew := tgbotapi.NewInlineKeyboardButtonData("➕ new", newArgs)
	btnEdit := tgbotapi.NewInlineKeyboardButtonData("✏ edit", editArgs)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(btnNew, btnEdit),
	)
}
