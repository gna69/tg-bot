package entity

// Object todo: redo interface, very bad
type Object interface {
	GetName() string
	GetId() uint
	GetCount() uint
	GetDescription() string
	GetUnit() string
	GetPrice() uint64
	GetCreatedAt() int64
	GetIngredients() string
	GetActions() string
	GetComplexity() uint8
	GetOwnerId() uint
	GetMembers() []string
	SetValue(step Step, value string) error
	ToString() string
}
