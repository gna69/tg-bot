package entity

type Action uint8

const (
	Nothing Action = iota
	ShowAll
	Add
	Change
	Delete
)
