package tg_bot

import (
	"context"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gna69/tg-bot/internal/adapters/pg"
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/gna69/tg-bot/internal/usecases"
)

type TgBot struct {
	api     *tgbotapi.BotAPI
	db      *pg.PostgresClient
	command *entity.Command
	stepper usecases.Stepper
}

func NewTelegramBot(token string, db *pg.PostgresClient) (*TgBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	command := entity.NewCommand()

	return &TgBot{
		api:     api,
		db:      db,
		command: command,
	}, nil
}

func (bot *TgBot) Run(ctx context.Context) error {
	chatCh, err := bot.getChat()
	if err != nil {
		return usecases.ErrGetChat
	}

	for message := range chatCh {
		if message.Message == nil {
			continue
		}

		if bot.isTextMessage(message.Message.Text) {
			chat := message.Message.Chat
			switch message.Message.Text {
			case entity.Start:
				bot.start(chat)
			case entity.Shopping:
				bot.shoppingMode(chat)
			case entity.Products:
				bot.productsMode(chat)
			case entity.Recipes:
				bot.recipesMode(chat)
			case entity.Workouts:
				bot.workoutsMode(chat)
			case entity.Stop:
				bot.stop(chat)
			default:
				bot.handle(ctx, message.Message)
			}
		}
	}

	return nil
}

func (bot *TgBot) getChat() (tgbotapi.UpdatesChannel, error) {
	chatSettings := tgbotapi.NewUpdate(DefaultOffset)
	chatSettings.Timeout = DefaultTimeout

	chat, err := bot.api.GetUpdatesChan(chatSettings)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (bot *TgBot) sendMessage(chatId int64, msg string) {
	_, _ = bot.api.Send(tgbotapi.NewMessage(chatId, msg))
}

func (bot *TgBot) isTextMessage(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.String && input.(string) != ""
}

func (bot *TgBot) isEnabled() bool {
	switch bot.command.GetCommand() {
	case entity.Stop:
		return false
	default:
		return true
	}
}
