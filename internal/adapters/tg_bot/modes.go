package tg_bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gna69/tg-bot/internal/adapters/stepper"
	"github.com/gna69/tg-bot/internal/usecases"
)

func (bot *TgBot) start(chat *tgbotapi.Chat) {
	if bot.isEnabled() {
		bot.sendMessage(chat.ID, "Я уже работаю!")
		return
	}
	bot.mode = Start
	bot.sendMessage(chat.ID, WelcomeMessage)
}

func (bot *TgBot) stop(chat *tgbotapi.Chat) {
	if !bot.isEnabled() {
		bot.sendMessage(chat.ID, "Я уже выключен!")
		return
	}
	bot.changeMode(Stop, FarewellMessage, chat)
}

func (bot *TgBot) shoppingMode(chat *tgbotapi.Chat) {
	shoppingStepper, err := stepper.New(usecases.ShoppingSteps)
	if err != nil {

	}

	bot.stepper = shoppingStepper
	bot.changeMode(Shopping, ShoppingBanner, chat)
}

func (bot *TgBot) productsMode(chat *tgbotapi.Chat) {
	bot.changeMode(Products, ProductsBanner, chat)
}

func (bot *TgBot) recipesMode(chat *tgbotapi.Chat) {
	bot.changeMode(Recipes, RecipesBanner, chat)
}

func (bot *TgBot) workoutsMode(chat *tgbotapi.Chat) {
	bot.changeMode(Workouts, WorkoutsBanner, chat)
}

func (bot *TgBot) changeMode(mode string, message string, chat *tgbotapi.Chat) {
	if !bot.isEnabled() {
		bot.sendMessage(chat.ID, AboutDisable)
		return
	}
	bot.disableChangesMode()
	bot.mode = mode
	bot.sendMessage(chat.ID, message)
}
