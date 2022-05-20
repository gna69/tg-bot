package stepper

import (
	"errors"
	"github.com/gna69/tg-bot/internal/entity"
)

type Stepper struct {
	changingStep bool
	currentStep  uint
	steps        []entity.Step
}

func New(steps []entity.Step) (*Stepper, error) {
	if len(steps) < 3 {
		return nil, errors.New("invalid count steps of stepper")
	}

	return &Stepper{
		steps: steps,
	}, nil
}

func (s *Stepper) Reset() {
	s.currentStep = 0
}

func (s *Stepper) NextStep() {
	s.currentStep += 1
	s.currentStep %= uint(len(s.steps))
}

func (s *Stepper) CurrentStep() entity.Step {
	return s.steps[s.currentStep]
}

func (s *Stepper) SetStep(newStep entity.Step) error {
	for stepNumber, step := range s.steps {
		if newStep == step {
			s.currentStep = uint(stepNumber)
			return nil
		}
	}

	return errors.New("Я не знаю таких данных, какие нужно изменить-то?")
}

func (s *Stepper) StepInfo() string {
	switch s.steps[s.currentStep] {
	case entity.Name:
		return "Введите название"
	case entity.Description:
		return "Введите описание"
	case entity.Count:
		return "Введите количество"
	case entity.Unit:
		return "Введите единицу измерения количества"
	case entity.Price:
		return "Введите цену"
	}
	return ""
}
