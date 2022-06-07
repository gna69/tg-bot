package main

import (
	"context"
	"github.com/gna69/tg-bot/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/gna69/tg-bot/internal/adapters/auth"
	"github.com/gna69/tg-bot/internal/adapters/pg"
	"github.com/gna69/tg-bot/internal/adapters/tg_bot"
	"github.com/gna69/tg-bot/internal/drivers/server"

	"github.com/jackc/pgx/v4"
)

var ctx = context.Background()

func main() {
	pgConn, err := pg.NewPgConnect(ctx)
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}

	defer func(pgConn *pgx.Conn, ctx context.Context) {
		err := pgConn.Close(ctx)
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}(pgConn, ctx)

	grpcConn, err := auth.NewGrpcConn()
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	defer func(grpcConn *grpc.ClientConn) {
		err := grpcConn.Close()
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}(grpcConn)

	authClient := proto.NewAuthServiceClient(grpcConn)

	tgBot, err := tg_bot.NewTelegramBot(cfg.BotToken, pgConn, authClient)
	if err != nil {
		log.Error().Msgf("Can't create telegram bot: ", err.Error())
		return
	}

	srv := server.New()
	err = srv.RunBots(ctx, tgBot)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
