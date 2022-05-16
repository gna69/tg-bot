package server

import (
	"github.com/gna69/tg-bot/internal/usecases"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) RunBots(bots ...usecases.Bot) error {
	for _, bot := range bots {
		err := bot.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
