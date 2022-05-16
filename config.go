package main

import (
	"github.com/caarlos0/env/v6"
	"log"
)

var cfg struct {
	BotToken string `env:"BOT_TOKEN"`
}

func init() {
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal("Error parsing config: ", err.Error())
	}
}
