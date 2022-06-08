package tg_bot

import (
	"context"
	"github.com/gna69/tg-bot/internal/entity"
	"strconv"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func (bot *TgBot) enableChangesMode(ctx context.Context, infoMsg string, chat *tgbotapi.Chat) bool {
	if !bot.context.stepper.IsChangingStep() {
		bot.sendMessage(chat.ID, infoMsg)
		bot.sendMessage(chat.ID, bot.showAll(ctx))
		bot.context.stepper.EnableChangingOption()
		return false
	}
	return true
}

func (bot *TgBot) disableChangesMode() {
	bot.context.stepper.Reset()
	bot.context.stepper.DisableChangingOption()
	bot.context.command.SetAction(entity.Nothing)
	bot.context.command.SetWorkingObjectId(0)
}

func (bot *TgBot) setObjectId(message *tgbotapi.Message) bool {
	if bot.context.command.GetWorkingObjectId() != 0 {
		return true
	}

	objId, err := strconv.Atoi(message.Text)
	if err != nil {
		bot.sendMessage(message.Chat.ID, "Я не понял, какой объект нужно изменить, повторишь?")
		return false
	}

	bot.context.command.SetWorkingObjectId(uint(objId))
	bot.sendMessage(message.Chat.ID, bot.context.stepper.UpdatingInfo())
	return false
}

func (bot *TgBot) setUpdatedStep(message *tgbotapi.Message) bool {
	if bot.context.stepper.CurrentStep() != entity.Waited {
		return true
	}

	err := bot.setUpdateInfo(message.Text)
	if err != nil {
		bot.sendMessage(message.Chat.ID, err.Error())
		return false
	}
	bot.sendMessage(message.Chat.ID, bot.context.stepper.StepInfo())
	return false
}
