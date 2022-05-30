package tg_bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gna69/tg-bot/internal/adapters/pg"
	"github.com/gna69/tg-bot/internal/adapters/stepper"
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/gna69/tg-bot/internal/usecases"
	"github.com/rs/zerolog/log"
)

func (bot *TgBot) start(chat *tgbotapi.Chat) {
	bot.command.SetCommand(entity.Start)
	bot.sendMessage(chat.ID, WelcomeMessage)
	log.Debug().Msg("Starting bot")
}

func (bot *TgBot) shoppingMode(chat *tgbotapi.Chat) {
	shoppingStepper, err := stepper.New(usecases.ShoppingSteps)
	if err != nil {
		return
	}

	bot.stepper = shoppingStepper
	bot.changeMode(entity.Shopping, ShoppingBanner, chat)
	log.Debug().Msgf("Setting mode to %s", entity.Shopping)
}

func (bot *TgBot) productsMode(chat *tgbotapi.Chat) {
	productsStepper, err := stepper.New(usecases.ProductsSteps)
	if err != nil {
		return
	}

	bot.stepper = productsStepper
	bot.changeMode(entity.Products, ProductsBanner, chat)
	log.Debug().Msgf("Setting mode to %s", entity.Products)
}

func (bot *TgBot) recipesMode(chat *tgbotapi.Chat) {
	recipesStepper, err := stepper.New(usecases.RecipesSteps)
	if err != nil {
		return
	}

	bot.stepper = recipesStepper
	bot.changeMode(entity.Recipes, RecipesBanner, chat)
	log.Debug().Msgf("Setting mode to %s", entity.Recipes)
}

func (bot *TgBot) workoutsMode(chat *tgbotapi.Chat) {
	bot.changeMode(entity.Workouts, WorkoutsBanner, chat)
}

func (bot *TgBot) changeMode(mode string, message string, chat *tgbotapi.Chat) {
	bot.disableChangesMode()
	bot.command.SetCommand(mode)
	bot.manager = pg.NewManager(mode, bot.db)
	bot.sendMessage(chat.ID, message)
}
