package main

import (
	"context"
	"github.com/gna69/tg-bot/internal/adapters/pg"
	"github.com/gna69/tg-bot/internal/adapters/tg_bot"
	"github.com/gna69/tg-bot/internal/drivers/server"
	"github.com/jackc/pgx/v4"
	"log"
)

var ctx = context.Background()

func main() {
	pgConn, err := pg.NewPgConnect(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func(pgConn *pgx.Conn, ctx context.Context) {
		err := pgConn.Close(ctx)
		if err != nil {
			log.Fatal("Error closing pg connect")
		}
	}(pgConn, ctx)

	tgBot, err := tg_bot.NewTelegramBot(cfg.BotToken, pgConn)
	if err != nil {
		log.Fatal("Can't create telegram bot: ", err.Error())
	}

	srv := server.New()
	err = srv.RunBots(ctx, tgBot)
	if err != nil {
		log.Fatal(err.Error())
	}
}
