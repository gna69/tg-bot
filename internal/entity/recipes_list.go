package entity

type Recipe struct {
	Id          uint
	Name        string
	Description string
	Ingredients string
	Actions     string
	Complexity  uint8
}
