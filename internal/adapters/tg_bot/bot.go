package tg_bot

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/gna69/tg-bot/internal/usecases"
)

type TgBot struct {
	api     *tgbotapi.BotAPI
	db      *pgx.Conn
	command *entity.Command
	manager usecases.Manager
	stepper usecases.Stepper
}

func NewTelegramBot(token string, db *pgx.Conn) (*TgBot, error) {
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

		log.Debug().
			Int("messageId", message.Message.MessageID).
			Str("messageText", message.Message.Text).
			Int("chatId", int(message.Message.Chat.ID)).
			Str("chatTitle", message.Message.Chat.Title).
			Str("from", message.Message.Chat.FirstName).Send()

		if message.Message.From.ID != 712226067 && message.Message.From.ID != 455932005 {
			bot.sendMessage(message.Message.Chat.ID, "Я нахожусь на этапе разработки, приходи в июле)")
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
