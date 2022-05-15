package entity

import "time"

type ShoppingList []Purchase

type Purchase struct {
	Id          uint
	Name        string
	Description string
	Count       uint8
	Unit        string
	Price       uint64
	CreatedAt   time.Time
}
