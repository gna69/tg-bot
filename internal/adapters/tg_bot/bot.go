package tg_bot

import (
	"context"
	"github.com/gna69/tg-bot/internal/adapters/synchronizer"
	"reflect"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/gna69/tg-bot/internal/usecases"
	pb "github.com/gna69/tg-bot/proto"

	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
)

type TgBot struct {
	api          *tgbotapi.BotAPI
	db           *pgx.Conn
	authCli      pb.AuthServiceClient
	synchronizer usecases.Synchronizer
	context      *botContext
}

type botContext struct {
	command    *entity.Command
	manager    usecases.Manager
	stepper    usecases.Stepper
	user       uint
	userGroups []int32
}

var userContext map[int]botContext

func NewTelegramBot(token string, db *pgx.Conn, authCli pb.AuthServiceClient) (*TgBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	command := entity.NewCommand()
	userContext = make(map[int]botContext)

	s := synchronizer.NewSynchronizer(db)

	return &TgBot{
		api: api,
		db:  db,
		context: &botContext{
			command: command,
		},
		synchronizer: s,
		authCli:      authCli,
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

		user := message.Message.From
		resp, err := bot.authCli.AuthUser(ctx, &pb.User{
			Id:           int32(user.ID),
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			UserName:     user.UserName,
			LanguageCode: user.LanguageCode,
			IsBot:        user.IsBot,
			ChatId:       uint64(message.Message.Chat.ID),
		})
		if err != nil {
			log.Debug().Msgf("error auth user: %s", err.Error())
			continue
		}

		// todo: to redis
		if uint(user.ID) != bot.context.user {
			if userContext, ok := userContext[user.ID]; ok {
				bot.context = &userContext
			} else {
				bot.context.command = entity.NewCommand()
			}
			bot.context.user = uint(user.ID)
		}

		bot.context.userGroups = resp.UserGroups
		err = bot.synchronizer.Synchronize(ctx, bot.context.userGroups, bot.context.user)
		if err != nil {
			log.Error().
				Str("synchronizer", "error synchronize user data").
				Err(err).
				Send()
		}

		log.Debug().
			Int("messageId", message.Message.MessageID).
			Str("messageText", message.Message.Text).
			Int("chatId", int(message.Message.Chat.ID)).
			Str("chatTitle", message.Message.Chat.Title).
			Str("from", message.Message.Chat.FirstName).Send()
		//todo: remove
		if user.ID != 712226067 && user.ID != 455932005 {
			bot.sendMessage(message.Message.Chat.ID, "Я нахожусь на этапе разработки, приходи в июле)")
			continue
		}

		if bot.isTextMessage(message.Message.Text) {
			chat := message.Message.Chat
			switch message.Message.Text {
			case entity.Start:
				bot.start(chat)
			case entity.Groups:
				bot.groupsMode(chat)
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

		userContext[user.ID] = *bot.context
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
