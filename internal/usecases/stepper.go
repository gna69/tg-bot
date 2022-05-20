package usecases

import "github.com/gna69/tg-bot/internal/entity"

var ShoppingSteps = []entity.Step{entity.Waited, entity.Name, entity.Count, entity.Description, entity.Unit, entity.Price, entity.End}

type Stepper interface {
	Reset()
	NextStep()
	CurrentStep() entity.Step
	IsChangingStep() bool
	EnableChangingOption()
	DisableChangingOption()
	SetStep(step entity.Step) error
	StepInfo() string
}
