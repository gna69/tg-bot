package usecases

import "github.com/gna69/tg-bot/internal/entity"

var ShoppingSteps = []entity.Step{entity.Waited, entity.Name, entity.Description, entity.Count, entity.Unit, entity.Price, entity.End}
var ProductsSteps = []entity.Step{entity.Waited, entity.Name, entity.Count, entity.End}
var RecipesSteps = []entity.Step{entity.Waited, entity.Name, entity.Description, entity.Ingredients, entity.Actions, entity.Complexity, entity.End}
var GroupsSteps = []entity.Step{entity.Waited, entity.Name, entity.Members, entity.End}

type Stepper interface {
	Reset()
	NextStep()
	CurrentStep() entity.Step
	IsChangingStep() bool
	EnableChangingOption()
	DisableChangingOption()
	SetStep(step uint) error
	StepInfo() string
	UpdatingInfo() string
}
