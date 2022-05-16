package main

import (
	"github.com/gna69/tg-bot/internal/adapters/tg_bot"
	"github.com/gna69/tg-bot/internal/drivers/server"
	"log"
)

func main() {
	tgBot, err := tg_bot.NewTelegramBot(cfg.BotToken)
	if err != nil {
		log.Fatal("Can't create telegram bot")
	}

	srv := server.New()
	err = srv.RunBots(tgBot)
	if err != nil {
		log.Fatal(err.Error())
	}
}
