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

	if !bot.enabled {
		bot.sendMessage(chatId, AboutDisable)
		return
	}

	if bot.context.action == Nothing {
		err := bot.setAction(message)
		if err != nil {
			bot.sendMessage(chatId, err.Error())
			return
		}
	}

	switch bot.context.action {
	case ShowAll:
		bot.sendMessage(chatId, bot.showAll(ctx))
		bot.context.action = Nothing
	case Add:
		if bot.context.step == Waited {
			bot.context.step = Name
			bot.sendMessage(chatId, StepInfoMessage(bot.context.step))
			return
		}

		err := bot.add(ctx, message.Text)
		if err != nil {
			bot.sendMessage(chatId, err.Error())
			return
		}

		bot.sendMessage(chatId, StepInfoMessage(bot.context.step))
	case Change:
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
	case Delete:
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

	if bot.context.step == Waited {
		bot.sendMessage(chatId, successInfo)
	}
}

func (bot *TgBot) setAction(message *tgbotapi.Message) error {
	newAction, err := strconv.Atoi(strings.ReplaceAll(message.Text, ".", ""))
	if err != nil {
		return usecases.ErrNoOption
	}

	bot.context.action = action(newAction)
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

	if bot.context.step == End {
		switch bot.mode {
		case Shopping:
			err := bot.db.ShoppingManager.AddPurchase(ctx, bot.context.purchase)
			if err != nil {
				return err
			}
			bot.context.purchase = &entity.Purchase{}
		}

		bot.context.step = Waited
		bot.context.action = Nothing
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

	switch step(updatingInfo) {
	case Name:
		bot.context.step = Name
	case Description:
		bot.context.step = Description
	case Count:
		bot.context.step = Count
	case Unit:
		bot.context.step = Unit
	case Price:
		bot.context.step = Price
	default:
		return errors.New("Я не знаю таких данных, какие нужно изменить-то?")
	}
	return nil
}

func (bot *TgBot) setInfo(value string) error {
	switch bot.context.step {
	case Name:
		if value == "" {
			return errors.New("не смог получить название, напиши еще разочек, пожалуйста")
		}
		bot.context.purchase.Name = value
		bot.context.step = Count
	case Count:
		count, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("Я не разобрался, подскажи, пожалуйста, какое количество необходимо?")
		}
		bot.context.purchase.Count = uint8(count)
		bot.context.step = Description
	case Description:
		bot.context.purchase.Description = value
		bot.context.step = Unit
	case Unit:
		bot.context.purchase.Unit = value
		bot.context.step = Price
	case Price:
		price, err := strconv.Atoi(strings.ReplaceAll(value, " ", ""))
		if err != nil {
			return errors.New("Что-то я не разобрался с ценой, можешь прислать еще раз?")
		}
		bot.context.purchase.Price = uint64(price)
		bot.context.purchase.CreatedAt = time.Now()
		bot.context.step = End
	}
	return nil
}
