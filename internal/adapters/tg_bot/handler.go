package tg_bot

import (
	"context"
	"errors"
	"github.com/gna69/tg-bot/internal/entity"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gna69/tg-bot/internal/usecases"
)

const (
	updatingInfo = "Что будем изменять?(просто пришли номер)"
	deletingInfo = "Что будем удалять?(просто пришли номер)"
	successInfo  = "Все прошло отлично, если что-то еще нужно сделать, выбери команду в меню!"
)

func (bot *TgBot) handle(ctx context.Context, message *tgbotapi.Message) {
	chatId := message.Chat.ID

	if !bot.isEnabled() {
		bot.sendMessage(chatId, AboutDisable)
		return
	}

	if bot.context.action == entity.Nothing {
		err := bot.setAction(message)
		if err != nil {
			bot.sendMessage(chatId, err.Error())
			return
		}
	}

	switch bot.context.action {
	case entity.ShowAll:
		bot.sendMessage(chatId, bot.showAll(ctx))
		bot.context.action = entity.Nothing
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

	bot.context.action = entity.Action(newAction)
	return nil
}

func (bot *TgBot) showAll(ctx context.Context) string {
	list := ""
	switch bot.mode {
	case Shopping:
		purchases, err := bot.db.ShoppingManager.GetPurchases(ctx)
		if err != nil {
			return "Не удалось получить список покупок, попробуй еще разок!"
		}
		list = bot.db.ShoppingManager.String(purchases)
	}

	if len(list) == 0 {
		return usecases.EmptyList
	}
	return list
}

func (bot *TgBot) add(ctx context.Context, message string) error {
	switch bot.mode {
	case Shopping:
		err := bot.setInfo(message)
		if err != nil {
			return err
		}
	}

	if bot.stepper.CurrentStep() == entity.End {
		switch bot.mode {
		case Shopping:
			err := bot.db.ShoppingManager.AddPurchase(ctx, bot.context.purchase)
			if err != nil {
				return err
			}
			bot.context.purchase = &entity.Purchase{}
		}

		bot.stepper.Reset()
		bot.context.action = entity.Nothing
	}

	return nil
}

func (bot *TgBot) update(ctx context.Context, message *tgbotapi.Message) error {
	switch bot.mode {
	case Shopping:
		purchase, err := bot.db.ShoppingManager.GetPurchase(ctx, bot.context.objectId)
		if err != nil {
			return err
		}
		bot.context.purchase = purchase

		err = bot.setInfo(message.Text)
		if err != nil {
			return err
		}

		err = bot.db.ShoppingManager.UpdatePurchase(ctx, bot.context.purchase)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bot *TgBot) delete(ctx context.Context, message *tgbotapi.Message) error {
	objId, err := strconv.Atoi(message.Text)
	if err != nil {
		return err
	}

	switch bot.mode {
	case Shopping:
		err = bot.db.ShoppingManager.DeletePurchase(ctx, uint(objId))
		if err != nil {
			return err
		}
	}
	return nil
}

func (bot *TgBot) setUpdateInfo(message string) error {
	updatingInfo, err := strconv.Atoi(message)
	if err != nil {
		return errors.New("Я не знаю таких данных, какие нужно изменить-то?")
	}

	err = bot.stepper.SetStep(entity.Step(updatingInfo))
	if err != nil {
		return err
	}
	return nil
}

func (bot *TgBot) setInfo(value string) error {
	switch bot.stepper.CurrentStep() {
	case entity.Name:
		if value == "" {
			return errors.New("Не смог получить название, напиши еще разочек, пожалуйста!")
		}
		bot.context.purchase.Name = value
	case entity.Count:
		count, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("Я не разобрался, подскажи, пожалуйста, какое количество необходимо?")
		}
		bot.context.purchase.Count = uint8(count)
	case entity.Description:
		bot.context.purchase.Description = value
	case entity.Unit:
		bot.context.purchase.Unit = value
	case entity.Price:
		price, err := strconv.Atoi(strings.ReplaceAll(value, " ", ""))
		if err != nil {
			return errors.New("Что-то я не разобрался с ценой, можешь прислать еще раз?")
		}
		bot.context.purchase.Price = uint64(price)
		bot.context.purchase.CreatedAt = time.Now()
	}

	bot.stepper.NextStep()
	return nil
}
