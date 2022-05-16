package tg_bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gna69/tg-bot/internal/usecases"
	"reflect"
)

const (
	DefaultOffset  = 0
	DefaultTimeout = 60
	WelcomeMessage = `Привет, я твой бот помощник по дому.
Я помогу тебе вести свой список покупок, продуктов, рецептов и тренировок.`
)

type TgBot struct {
	api *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (*TgBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TgBot{api: api}, nil
}

func (bot *TgBot) Run() error {
	chat, err := bot.getChat()
	if err != nil {
		return usecases.ErrGetChat
	}

	for message := range *chat {
		if message.Message == nil {
			continue
		}

		if bot.string(message.Message.Text) {
			switch message.Message.Text {
			case "/start":
				bot.sendMessage(message.Message.Chat.ID, WelcomeMessage)
			case "/shopping":

			default:

			}
		}
	}

	return nil
}

func (bot *TgBot) getChat() (*tgbotapi.UpdatesChannel, error) {
	chatSettings := tgbotapi.NewUpdate(DefaultOffset)
	chatSettings.Timeout = DefaultTimeout

	chat, err := bot.api.GetUpdatesChan(chatSettings)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (bot *TgBot) sendMessage(chatId int64, msg string) {
	_, _ = bot.api.Send(tgbotapi.NewMessage(chatId, msg))
}

func (bot *TgBot) string(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.String && input.(string) != ""
}
