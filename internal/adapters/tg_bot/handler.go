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

	if bot.command.GetAction() == entity.Nothing {
		err := bot.setAction(message)
		if err != nil {
			bot.sendMessage(chatId, err.Error())
			return
		}
	}

	switch bot.command.GetAction() {
	case entity.ShowAll:
		bot.sendMessage(chatId, bot.showAll(ctx))
		bot.command.SetAction(entity.Nothing)
	case entity.Add:
		if bot.stepper.CurrentStep() == entity.Waited {
			bot.stepper.NextStep()
			bot.sendMessage(chatId, bot.stepper.StepInfo())
			return
		}

		err := bot.add(ctx, message.Text)
		if err != nil {
			bot.sendMessage(chatId, err.Error())
			return
		}

		bot.sendMessage(chatId, bot.stepper.StepInfo())
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
	default:
		bot.sendMessage(chatId, usecases.ErrNoOption.Error())
		return
	}

	if bot.stepper.CurrentStep() == entity.Waited {
		bot.sendMessage(chatId, successInfo)
	}
}

func (bot *TgBot) setAction(message *tgbotapi.Message) error {
	newAction, err := strconv.Atoi(strings.ReplaceAll(message.Text, ".", ""))
	if err != nil {
		return usecases.ErrNoOption
	}

	bot.command.SetAction(entity.Action(newAction))
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
	result, err := bot.manager.GetAll(ctx)
	if err != nil {
		return err.Error()
	}

	return toString(result)
}

func (bot *TgBot) add(ctx context.Context, message string) error {
	err := bot.command.Object.SetValue(bot.stepper.CurrentStep(), message)
	if err != nil {
		return err
	}
	bot.stepper.NextStep()

	if bot.stepper.CurrentStep() == entity.End {
		err := bot.manager.Add(ctx, bot.command.Object)
		if err != nil {
			return err
		}

		bot.command.Object = nil
		bot.stepper.Reset()
		bot.command.SetAction(entity.Nothing)
	}
	return nil
}

func (bot *TgBot) update(ctx context.Context, message *tgbotapi.Message) error {
	updatingObject, err := bot.manager.Get(ctx, bot.command.GetWorkingObjectId())
	if err != nil {
		return err
	}
	bot.command.Object = updatingObject

	err = bot.command.Object.SetValue(bot.stepper.CurrentStep(), message.Text)
	if err != nil {
		return err
	}

	err = bot.manager.Update(ctx, bot.command.Object)
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

	err = bot.manager.Delete(ctx, uint(objId))
	return err
}

func (bot *TgBot) setUpdateInfo(message string) error {
	updatingInfo, err := strconv.Atoi(message)
	if err != nil {
		return errors.New("Я не знаю таких данных, какие нужно изменить-то?")
	}

	err = bot.stepper.SetStep(uint(updatingInfo))
	if err != nil {
		return err
	}
	return nil
}
