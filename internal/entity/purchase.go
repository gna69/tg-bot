package entity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Purchase struct {
	Id          uint
	Name        string
	Description string
	Count       uint8
	Unit        string
	Price       uint64
	CreatedAt   time.Time
}

func (p *Purchase) ToString() string {
	strView := fmt.Sprintf("\n%d) %s %d %s %d руб.", p.Id, p.Name, p.Count, p.Unit, p.Price)
	if len(p.Description) != 0 {
		strView += fmt.Sprintf("\n%s", p.Description)
	}
	strView += fmt.Sprintf("\n%s\n", p.CreatedAt.Format("15:04 02:January:2006"))
	return strView
}

func (p *Purchase) SetValue(step Step, value string) error {
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
		p.Count = uint8(count)
	case Description:
		p.Description = value
	case Unit:
		p.Unit = value
	case Price:
		price, err := strconv.Atoi(strings.ReplaceAll(value, " ", ""))
		if err != nil {
			return errors.New("Что-то я не разобрался с ценой, можешь прислать еще раз?")
		}
		p.Price = uint64(price)
		p.CreatedAt = time.Now()
	}
	return nil
}
