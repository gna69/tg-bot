package tg_bot

import (
	"github.com/gna69/tg-bot/internal/adapters/pg"
	"github.com/gna69/tg-bot/internal/adapters/stepper"
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/gna69/tg-bot/internal/usecases"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/rs/zerolog/log"
)

func (bot *TgBot) start(chat *tgbotapi.Chat) {
	bot.context.command.SetCommand(entity.Start)
	bot.sendMessage(chat.ID, WelcomeMessage)
	log.Debug().Msg("Starting bot")
}

func (bot *TgBot) groupsMode(chat *tgbotapi.Chat) {
	groupsStepper, err := stepper.New(usecases.GroupsSteps)
	if err != nil {
		return
	}

	bot.context.stepper = groupsStepper
	bot.context.command.Object = &entity.Group{OwnerId: bot.context.user}
	bot.changeMode(entity.Groups, GroupsBanner, chat)
	logChangeMode(entity.Groups)
}

func (bot *TgBot) shoppingMode(chat *tgbotapi.Chat) {
	shoppingStepper, err := stepper.New(usecases.ShoppingSteps)
	if err != nil {
		return
	}

	bot.context.stepper = shoppingStepper
	bot.context.command.Object = &entity.Purchase{OwnerId: bot.context.user}
	bot.changeMode(entity.Shopping, ShoppingBanner, chat)
	logChangeMode(entity.Shopping)
}

func (bot *TgBot) productsMode(chat *tgbotapi.Chat) {
	productsStepper, err := stepper.New(usecases.ProductsSteps)
	if err != nil {
		return
	}

	bot.context.stepper = productsStepper
	bot.context.command.Object = &entity.Product{OwnerId: bot.context.user}
	bot.changeMode(entity.Products, ProductsBanner, chat)
	logChangeMode(entity.Products)
}

func (bot *TgBot) recipesMode(chat *tgbotapi.Chat) {
	recipesStepper, err := stepper.New(usecases.RecipesSteps)
	if err != nil {
		return
	}

	bot.context.stepper = recipesStepper
	bot.context.command.Object = &entity.Recipe{OwnerId: bot.context.user}
	bot.changeMode(entity.Recipes, RecipesBanner, chat)
	logChangeMode(entity.Recipes)
}

func (bot *TgBot) workoutsMode(chat *tgbotapi.Chat) {
	bot.changeMode(entity.Workouts, WorkoutsBanner, chat)
}

func (bot *TgBot) changeMode(mode string, message string, chat *tgbotapi.Chat) {
	bot.disableChangesMode()
	bot.context.command.SetCommand(mode)
	bot.context.manager = pg.NewManager(mode, bot.db, bot.authCli)
	bot.sendMessage(chat.ID, message)
}

func logChangeMode(mode string) {
	log.Debug().Msgf("Setting mode to %s", mode)
}
