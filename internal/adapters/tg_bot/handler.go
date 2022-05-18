package tg_bot

import (
	"context"
	"strconv"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gna69/tg-bot/internal/usecases"
)

func (bot *TgBot) handle(ctx context.Context, message *tgbotapi.Message) {
	if bot.context.operation == Nothing {
		err := bot.setOperation(message)
		if err != nil {
			bot.sendMessage(message.Chat.ID, err.Error())
			return
		}
	}

	switch bot.context.operation {
	case ShowAll:
		bot.sendMessage(message.Chat.ID, bot.showAll(ctx))
		bot.context.operation = Nothing
	case Add:
	case Change:
	case Delete:
	default:
		bot.sendMessage(message.Chat.ID, usecases.ErrNoOption.Error())
	}

}

func (bot *TgBot) setOperation(message *tgbotapi.Message) error {
	op, err := strconv.Atoi(strings.ReplaceAll(message.Text, ".", ""))
	if err != nil {
		return usecases.ErrNoOption
	}
	bot.context.operation = operation(op)
	return nil
}

func (bot *TgBot) showAll(ctx context.Context) string {
	list := ""
	switch bot.mode {
	case Shopping:
		purchases := bot.db.ShoppingManager.GetPurchases(ctx)
		list = bot.db.ShoppingManager.String(purchases)
	}

	if len(list) == 0 {
		return usecases.EmptyList
	}
	return list
}
