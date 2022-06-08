package tg_bot

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/gna69/tg-bot/internal/entity"
	"github.com/gna69/tg-bot/internal/usecases"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

const (
	updatingInfo = "Что будем изменять?(просто пришли номер)"
	deletingInfo = "Что будем удалять?(просто пришли номер)"
	successInfo  = "Все прошло отлично, если что-то еще нужно сделать, выбери команду в меню!"
)

func (bot *TgBot) handle(ctx context.Context, message *tgbotapi.Message) {
	chatId := message.Chat.ID

	if bot.context.command.GetCommand() == entity.Start {
		return
	}

	if bot.context.command.GetAction() == entity.Nothing {
		err := bot.setAction(message)
		if err != nil {
			bot.sendMessage(chatId, err.Error())
			return
		}
	}

	switch bot.context.command.GetAction() {
	case entity.ShowAll:
		bot.sendMessage(chatId, bot.showAll(ctx))
		bot.context.command.SetAction(entity.Nothing)
	case entity.Add:
		if bot.context.stepper.CurrentStep() == entity.Waited {
			bot.context.stepper.NextStep()
			bot.sendMessage(chatId, bot.context.stepper.StepInfo())
			return
		}

		err := bot.add(ctx, message.Text)
		if err != nil {
			bot.sendMessage(chatId, err.Error())
			return
		}

		bot.sendMessage(chatId, bot.context.stepper.StepInfo())
	case entity.Change:
		if !bot.enableChangesMode(ctx, updatingInfo, message.Chat) {
			return
		}
		if !bot.setObjectId(message) {
			return
		}
		if !bot.setUpdatedStep(message) {
			return
		}

		err := bot.update(ctx, message)
		if err != nil {
			bot.sendMessage(chatId, err.Error())
			return
		}

		bot.disableChangesMode()
	case entity.Delete:
		if !bot.enableChangesMode(ctx, deletingInfo, message.Chat) {
			return
		}

		err := bot.delete(ctx, message)
		if err != nil {
			bot.sendMessage(chatId, err.Error())
		}

		bot.disableChangesMode()
	case entity.RemoveMember:
		fallthrough
	default:
		bot.sendMessage(chatId, usecases.ErrNoOption.Error())
		return
	}

	if bot.context.stepper.CurrentStep() == entity.Waited {
		bot.sendMessage(chatId, successInfo)
	}
}

func (bot *TgBot) setAction(message *tgbotapi.Message) error {
	newAction, err := strconv.Atoi(strings.ReplaceAll(message.Text, ".", ""))
	if err != nil {
		return usecases.ErrNoOption
	}

	bot.context.command.SetAction(entity.Action(newAction))
	return nil
}

func toString(objs []entity.Object) string {
	str := ""
	for _, obj := range objs {
		str += obj.ToString()
	}
	if len(str) == 0 {
		return usecases.EmptyList
	}
	return str
}

func (bot *TgBot) showAll(ctx context.Context) string {
	result, err := bot.context.manager.GetAll(ctx, bot.context.user, bot.context.userGroups)
	if err != nil {
		return err.Error()
	}

	return toString(result)
}

func (bot *TgBot) add(ctx context.Context, message string) error {
	err := bot.context.command.Object.SetValue(bot.context.stepper.CurrentStep(), message)
	if err != nil {
		return err
	}
	bot.context.stepper.NextStep()

	if bot.context.stepper.CurrentStep() == entity.End {
		err := bot.context.manager.Add(ctx, bot.context.command.Object)
		if err != nil {
			return err
		}

		bot.context.command.Object = nil
		bot.context.stepper.Reset()
		bot.context.command.SetAction(entity.Nothing)
	}
	return nil
}

func (bot *TgBot) update(ctx context.Context, message *tgbotapi.Message) error {
	updatingObject, err := bot.context.manager.Get(ctx, bot.context.command.GetWorkingObjectId(), bot.context.user, bot.context.userGroups)
	if err != nil {
		return err
	}
	bot.context.command.Object = updatingObject

	err = bot.context.command.Object.SetValue(bot.context.stepper.CurrentStep(), message.Text)
	if err != nil {
		return err
	}

	err = bot.context.manager.Update(ctx, bot.context.command.Object)
	if err != nil {
		return err
	}

	return nil
}

func (bot *TgBot) delete(ctx context.Context, message *tgbotapi.Message) error {
	objId, err := strconv.Atoi(message.Text)
	if err != nil {
		return err
	}

	err = bot.context.manager.Delete(ctx, uint(objId))
	return err
}

func (bot *TgBot) setUpdateInfo(message string) error {
	updatingInfo, err := strconv.Atoi(message)
	if err != nil {
		return errors.New("Я не знаю таких данных, какие нужно изменить-то?")
	}

	err = bot.context.stepper.SetStep(uint(updatingInfo))
	if err != nil {
		return err
	}
	return nil
}
