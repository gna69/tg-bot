package entity

import (
	"errors"
	"fmt"
	"strconv"
)

type Product struct {
	Id         uint
	Name       string
	TotalCount uint8
	OwnerId    uint
	Groups     []int32
}

func (p *Product) ToString() string {
	return fmt.Sprintf("\n%d) %s %d шт.", p.Id, p.Name, p.TotalCount)
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetId() uint {
	return p.Id
}

func (p *Product) GetCount() uint {
	return uint(p.TotalCount)
}

func (p *Product) GetOwnerId() uint {
	return p.OwnerId
}

func (p *Product) SetOwnerId(ownerId uint) {
	p.OwnerId = ownerId
}

func (p *Product) GetDescription() string {
	return ""
}

func (p *Product) GetUnit() string {
	return ""
}

func (p *Product) GetPrice() uint64 {
	return 0
}

func (p *Product) GetCreatedAt() int64 {
	return 0
}

func (p *Product) GetMembers() []string {
	return nil
}

func (p *Product) GetIngredients() string {
	return ""
}

func (p *Product) GetActions() string {
	return ""
}

func (p *Product) GetComplexity() uint8 {
	return 0
}

func (p *Product) SetValue(step Step, value string) error {
	switch step {
	case Name:
		if value == "" {
			return errors.New("Не смог получить название, напиши еще разочек, пожалуйста!")
		}
		p.Name = value
	case Count:
		count, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("Я не разобрался, подскажи, пожалуйста, какое количество необходимо?")
		}
		p.TotalCount = uint8(count)
	}
	return nil
}
