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
	api      *tgbotapi.BotAPI
	db       *pg.PostgresClient
	enabled  bool
	mode     string
	context  modeContext
	entities entities
}

type modeContext struct {
	action   action
	changes  bool
	objectId uint
	step     step
	purchase *entity.Purchase
}

type entities struct {
	purchase *entity.Purchase
	product  *entity.Product
	recipe   *entity.Recipe
	workout  *entity.Workout
}

func NewTelegramBot(token string, db *pg.PostgresClient) (*TgBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TgBot{
		api:     api,
		db:      db,
		enabled: false,
		mode:    Stop,
		context: modeContext{
			action:   Nothing,
			changes:  false,
			objectId: 0,
			step:     Waited,
			purchase: &entity.Purchase{},
		},
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
			case Start:
				bot.start(chat)
			case Shopping:
				bot.shoppingMode(chat)
			case Products:
				bot.productsMode(chat)
			case Recipes:
				bot.recipesMode(chat)
			case Workouts:
				bot.workoutsMode(chat)
			case Stop:
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
