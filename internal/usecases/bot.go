package usecases

import (
	"context"
	"errors"
)

var (
	ErrGetChat  = errors.New("cant get bot chat")
	ErrNoOption = errors.New("Такой опции у меня нет!")
)

type Bot interface {
	Run(ctx context.Context) error
}
