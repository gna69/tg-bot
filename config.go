package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var cfg struct {
	InfoLevel string `env:"INFO_LEVEL" envDefault:"DEBUG"`
	BotToken  string `env:"BOT_TOKEN"`
}

func init() {
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error parsing config: %s", err.Error())
	}

	switch cfg.InfoLevel {
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

}
