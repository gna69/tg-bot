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
}

func (p *Product) ToString() string {
	return fmt.Sprintf("\n%d) %s %d шт.", p.Id, p.Name, p.TotalCount)
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
