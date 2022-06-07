package entity

import (
	"errors"
	"fmt"
	"strconv"
)

type Recipe struct {
	Id          uint
	Name        string
	Description string
	Ingredients string
	Actions     string
	Complexity  uint8
	OwnerId     uint
}

func (r *Recipe) ToString() string {
	strView := fmt.Sprintf("\n%d) %s", r.Id, r.Name)
	if len(r.Description) != 0 {
		strView += fmt.Sprintf("\n%s", r.Description)
	}
	strView += fmt.Sprintf("\n\nИнгридиенты:\n%s", r.Ingredients)
	strView += fmt.Sprintf("\n\n%s", r.Actions)
	strView += fmt.Sprintf("\nСложность: %d/5\n", r.Complexity)
	return strView
}

func (r *Recipe) GetName() string {
	return r.Name
}

func (r *Recipe) GetIngredients() string {
	return r.Ingredients
}

func (r *Recipe) GetActions() string {
	return r.Actions
}

func (r *Recipe) GetComplexity() uint8 {
	return r.Complexity
}

func (r *Recipe) GetId() uint {
	return r.Id
}

func (r *Recipe) GetDescription() string {
	return r.Description
}

func (r *Recipe) GetOwnerId() uint {
	return r.OwnerId
}

func (r *Recipe) SetOwnerId(ownerId uint) {
	r.OwnerId = ownerId
}

func (r *Recipe) SetValue(step Step, value string) error {
	switch step {
	case Name:
		if value == "" {
			return errors.New("Не смог получить название, напиши еще разочек, пожалуйста!")
		}
		r.Name = value
	case Description:
		r.Description = value
	case Ingredients:
		r.Ingredients = value
	case Actions:
		r.Actions = value
	case Complexity:
		complexity, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("Я не разобрался, подскажи, пожалуйста, какая сложность у рецепта?")
		}
		if complexity < 0 || complexity > 5 {
			return errors.New("У рецептлв 5ти бальная шкала сложности, попробуй еще разочек!")
		}
		r.Complexity = uint8(complexity)
	}
	return nil
}

func (r *Recipe) GetUnit() string {
	return ""
}

func (r *Recipe) GetPrice() uint64 {
	return 0
}

func (r *Recipe) GetCreatedAt() int64 {
	return 0
}

func (r *Recipe) GetCount() uint {
	return 0
}

func (r *Recipe) GetMembers() []string {
	return nil
}
