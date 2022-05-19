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
		if bot.context.additionalInfo == Waited {
			bot.context.additionalInfo = Name
			bot.sendMessage(message.Chat.ID, AddedInfoMessage(bot.context.additionalInfo))
			return
		}

		err := bot.add(ctx, message.Text)
		if err != nil {
			bot.sendMessage(message.Chat.ID, err.Error())
			return
		}

		bot.sendMessage(message.Chat.ID, AddedInfoMessage(bot.context.additionalInfo))
	case Change:
		if !bot.context.updating {
			bot.sendMessage(message.Chat.ID, UpdatingList)
			bot.context.updating = true
			return
		}

		if bot.context.additionalInfo == Waited {
			err := bot.setUpdateInfo(message.Text)
			if err != nil {
				bot.sendMessage(message.Chat.ID, err.Error())
				return
			}
			bot.sendMessage(message.Chat.ID, AddedInfoMessage(bot.context.additionalInfo))
			return
		}

		err := bot.update(ctx, message)
		if err != nil {
			bot.sendMessage(message.Chat.ID, err.Error())
			return
		}

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

	if bot.context.additionalInfo == End {
		switch bot.mode {
		case Shopping:
			err := bot.db.ShoppingManager.AddPurchase(ctx, bot.context.purchase)
			if err != nil {
				return err
			}
			bot.context.purchase = &entity.Purchase{}
		}

		bot.context.additionalInfo = Waited
	}

	return nil
}

func (bot *TgBot) update(ctx context.Context, message *tgbotapi.Message) error {
	switch bot.mode {
	case Shopping:
		err := bot.setInfo(message.Text)
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

func (bot *TgBot) setUpdateInfo(message string) error {
	updatingInfo, err := strconv.Atoi(message)
	if err != nil {
		return errors.New("Я не знаю таких данных, какие нужно изменить-то?")
	}

	switch addedInfo(updatingInfo) {
	case Name:
		bot.context.additionalInfo = Name
	case Description:
		bot.context.additionalInfo = Description
	case Count:
		bot.context.additionalInfo = Count
	case Unit:
		bot.context.additionalInfo = Unit
	case Price:
		bot.context.additionalInfo = Price
	default:
		return errors.New("Я не знаю таких данных, какие нужно изменить-то?")
	}
	return nil
}

func (bot *TgBot) setInfo(value string) error {
	switch bot.context.additionalInfo {
	case Name:
		if value == "" {
			return errors.New("не смог получить название, напиши еще разочек, пожалуйста")
		}
		bot.context.purchase.Name = value
		bot.context.additionalInfo = Count
	case Count:
		count, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("Я не разобрался, подскажи, пожалуйста, какое количество необходимо?")
		}
		bot.context.purchase.Count = uint8(count)
		bot.context.additionalInfo = Description
	case Description:
		bot.context.purchase.Description = value
		bot.context.additionalInfo = Unit
	case Unit:
		bot.context.purchase.Unit = value
		bot.context.additionalInfo = Price
	case Price:
		price, err := strconv.Atoi(strings.ReplaceAll(value, " ", ""))
		if err != nil {
			return errors.New("Что-то я не разобрался с ценой, можешь прислать еще раз?")
		}
		bot.context.purchase.Price = uint64(price)
		bot.context.purchase.CreatedAt = time.Now()
		bot.context.additionalInfo = End
	}
	return nil
}
