package entity

import (
	"fmt"
	"strings"
)

type Group struct {
	Id      uint
	OwnerId uint
	Name    string
	Members []string
}

func (g *Group) GetName() string {
	return g.Name
}

func (g *Group) GetId() uint {
	return g.Id
}

func (g *Group) GetOwnerId() uint {
	return g.OwnerId
}

func (g *Group) SetOwnerId(ownerId uint) {
	g.OwnerId = ownerId
}

func (g *Group) GetMembers() []string {
	return g.Members
}

func (g *Group) SetValue(step Step, value string) error {
	switch step {
	case Name:
		g.Name = value
	case Members:
		value = strings.ReplaceAll(value, " ", "")
		members := strings.Split(value, ",")
		if len(members) == 0 {
			return fmt.Errorf("Список членов группы пустой!")
		}
		g.Members = members
	}
	return nil
}

func (g *Group) ToString() string {
	strView := fmt.Sprintf("\n%d) Группа %s", g.Id, g.Name)
	strView += "\nСписок участников:"
	for idx, member := range g.Members {
		strView += fmt.Sprintf("\n\t\t%d) %s", idx+1, member)
	}
	strView += "\n"
	return strView
}

func (g *Group) GetCount() uint {
	return 0
}

func (g *Group) GetDescription() string {
	return ""
}

func (g *Group) GetUnit() string {
	return ""
}

func (g *Group) GetPrice() uint64 {
	return 0
}

func (g *Group) GetCreatedAt() int64 {
	return 0
}

func (g *Group) GetIngredients() string {
	return ""
}

func (g *Group) GetActions() string {
	return ""
}

func (g *Group) GetComplexity() uint8 {
	return 0
}
