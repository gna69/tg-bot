package tg_bot

import (
	"context"
	"github.com/gna69/tg-bot/internal/entity"
	"strconv"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (bot *TgBot) enableChangesMode(ctx context.Context, infoMsg string, chat *tgbotapi.Chat) bool {
	if !bot.context.changes {
		bot.sendMessage(chat.ID, infoMsg)
		bot.sendMessage(chat.ID, bot.showAll(ctx))
		bot.context.changes = true
		return false
	}
	return true
}

func (bot *TgBot) disableChangesMode() {
	bot.stepper.Reset()
	bot.context.action = entity.Nothing
	bot.context.changes = false
	bot.context.objectId = 0
}

func (bot *TgBot) setObjectId(message *tgbotapi.Message) bool {
	if bot.context.objectId != 0 {
		return true
	}

	objId, err := strconv.Atoi(message.Text)
	if err != nil {
		bot.sendMessage(message.Chat.ID, "Я не понял, какой объект нужно изменить, повторишь?")
		return false
	}

	bot.context.objectId = uint(objId)
	bot.sendMessage(message.Chat.ID, UpdatingList)
	return false
}

func (bot *TgBot) setUpdatedStep(message *tgbotapi.Message) bool {
	if bot.stepper.CurrentStep() != entity.Waited {
		return true
	}

	err := bot.setUpdateInfo(message.Text)
	if err != nil {
		bot.sendMessage(message.Chat.ID, err.Error())
		return false
	}
	bot.sendMessage(message.Chat.ID, bot.stepper.StepInfo())
	return false
}
