package tg_bot

import tgbotapi "github.com/Syfaro/telegram-bot-api"

func (bot *TgBot) start(chat *tgbotapi.Chat) {
	if bot.enabled {
		bot.sendMessage(chat.ID, "Я уже работаю!")
		return
	}
	bot.enabled = true
	bot.changeMode(Stop, WelcomeMessage, chat)
}

func (bot *TgBot) stop(chat *tgbotapi.Chat) {
	if !bot.enabled {
		bot.sendMessage(chat.ID, "Я уже выключен!")
		return
	}
	bot.changeMode(Stop, FarewellMessage, chat)
	bot.enabled = false
}

func (bot *TgBot) shoppingMode(chat *tgbotapi.Chat) {
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
	if !bot.enabled {
		bot.sendMessage(chat.ID, AboutDisable)
		return
	}
	bot.currentMode = mode
	bot.sendMessage(chat.ID, message)
}
