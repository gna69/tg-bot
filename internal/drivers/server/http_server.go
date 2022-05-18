package server

import (
	"context"
	"fmt"
	"sync"

	"github.com/gna69/tg-bot/internal/usecases"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) RunBots(ctx context.Context, bots ...usecases.Bot) error {
	wg := sync.WaitGroup{}

	for _, bot := range bots {

		go func(bot usecases.Bot) {
			err := bot.Run(ctx)
			if err != nil {
				fmt.Println("err: ", err.Error())
			}

			wg.Done()
		}(bot)

		wg.Add(1)
	}

	fmt.Println("All bots are running ")
	wg.Wait()
	return nil
}