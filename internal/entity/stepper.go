package entity

type Step uint8

const (
	Waited Step = iota
	Name
	Description
	Count
	Unit
	Price
	Ingredients
	Actions
	Complexity
	Members
	End
)
