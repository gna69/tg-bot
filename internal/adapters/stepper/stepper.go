package stepper

import (
	"errors"
	"fmt"
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

func (s *Stepper) IsChangingStep() bool {
	return s.changingStep
}

func (s *Stepper) EnableChangingOption() {
	s.changingStep = true
}

func (s *Stepper) DisableChangingOption() {
	s.changingStep = false
}

func (s *Stepper) SetStep(newStep uint) error {
	if newStep < 1 || newStep >= uint(len(s.steps)-1) {
		return errors.New("Эту информацию я обновить не могу!")
	}
	s.currentStep = newStep
	return nil
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
	case entity.Ingredients:
		return "Введите ингридиенты (через дефис или двоеточие можно указать количество)"
	case entity.Actions:
		return "Введите последовательность действий"
	case entity.Complexity:
		return "Введите сложность рецепта по 5ти бальной шкале"
	}
	return ""
}

func (s *Stepper) UpdatingInfo() string {
	info := "Какую информацию хочешь поменять?\n"
	for idx, step := range s.steps {
		switch step {
		case entity.Name:
			info += fmt.Sprintf("%d) Название.\n", idx)
		case entity.Description:
			info += fmt.Sprintf("%d) Описание.\n", idx)
		case entity.Count:
			info += fmt.Sprintf("%d) Количество.\n", idx)
		case entity.Unit:
			info += fmt.Sprintf("%d) Еденицу измерения.\n", idx)
		case entity.Price:
			info += fmt.Sprintf("%d) Цену.\n", idx)
		}
	}
	return info
}
