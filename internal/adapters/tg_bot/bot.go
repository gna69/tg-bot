package tg_bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gna69/tg-bot/internal/usecases"
	"reflect"
)

type TgBot struct {
	api         *tgbotapi.BotAPI
	enabled     bool
	currentMode string
}

type modeStatus struct {
}

func NewTelegramBot(token string) (*TgBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TgBot{
		api:         api,
		enabled:     false,
		currentMode: Stop,
	}, nil
}

func (bot *TgBot) Run() error {
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
				bot.handle(message.Message)
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
