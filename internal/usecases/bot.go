package usecases

import "errors"

var (
	ErrGetChat = errors.New("cant get bot chat")
)

type Bot interface {
	Run() error
}
