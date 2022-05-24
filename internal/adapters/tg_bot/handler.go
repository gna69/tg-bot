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

	if !bot.isEnabled() {
		bot.sendMessage(chatId, AboutDisable)
		return
	}

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

func (bot *TgBot) showAll(ctx context.Context) string {
	list := ""
	switch bot.command.GetCommand() {
	case entity.Shopping:
		purchases, err := bot.db.ShoppingManager.GetPurchases(ctx)
		if err != nil {
			return "Не удалось получить список покупок, попробуй еще разок!"
		}
		list = bot.db.ShoppingManager.String(purchases)
	case entity.Products:
		products, err := bot.db.ProductsManager.GetProducts(ctx)
		if err != nil {
			return "Не удалось получить список покупок, попробуй еще разок!"
		}
		list = bot.db.ProductsManager.String(products)
	}

	if len(list) == 0 {
		return usecases.EmptyList
	}
	return list
}

func (bot *TgBot) add(ctx context.Context, message string) error {
	switch bot.command.GetCommand() {
	case entity.Shopping:
		err := bot.command.SetObjectValue(bot.stepper.CurrentStep(), message)
		if err != nil {
			return err
		}
	case entity.Products:
		err := bot.command.SetObjectValue(bot.stepper.CurrentStep(), message)
		if err != nil {
			return err
		}
	}
	bot.stepper.NextStep()

	if bot.stepper.CurrentStep() == entity.End {
		switch bot.command.GetCommand() {
		case entity.Shopping:
			err := bot.db.ShoppingManager.AddPurchase(ctx, bot.command.WorkingObject.Purchase)
			if err != nil {
				return err
			}
			bot.command.WorkingObject.Purchase = &entity.Purchase{}
		case entity.Products:
			err := bot.db.ProductsManager.AddProduct(ctx, bot.command.WorkingObject.Product)
			if err != nil {
				return err
			}
			bot.command.WorkingObject.Product = &entity.Product{}
		}

		bot.stepper.Reset()
		bot.command.SetAction(entity.Nothing)
	}
	return nil
}

func (bot *TgBot) update(ctx context.Context, message *tgbotapi.Message) error {
	switch bot.command.GetCommand() {
	case entity.Shopping:
		purchase, err := bot.db.ShoppingManager.GetPurchase(ctx, bot.command.GetWorkingObjectId())
		if err != nil {
			return err
		}
		bot.command.WorkingObject.Purchase = purchase

		err = bot.command.SetObjectValue(bot.stepper.CurrentStep(), message.Text)
		if err != nil {
			return err
		}

		err = bot.db.ShoppingManager.UpdatePurchase(ctx, bot.command.WorkingObject.Purchase)
		if err != nil {
			return err
		}
	case entity.Products:
		product, err := bot.db.ProductsManager.GetProduct(ctx, bot.command.GetWorkingObjectId())
		if err != nil {
			return err
		}
		bot.command.WorkingObject.Product = product

		err = bot.command.SetObjectValue(bot.stepper.CurrentStep(), message.Text)
		if err != nil {
			return err
		}

		err = bot.db.ProductsManager.UpdateProduct(ctx, bot.command.WorkingObject.Product)
	}

	return nil
}

func (bot *TgBot) delete(ctx context.Context, message *tgbotapi.Message) error {
	objId, err := strconv.Atoi(message.Text)
	if err != nil {
		return err
	}

	switch bot.command.GetCommand() {
	case entity.Shopping:
		err = bot.db.ShoppingManager.DeletePurchase(ctx, uint(objId))
		if err != nil {
			return err
		}
	case entity.Products:
		err = bot.db.ProductsManager.DeleteProduct(ctx, uint(objId))

	}
	return nil
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
