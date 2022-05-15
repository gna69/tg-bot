package managers

import (
	"github.com/gna69/tg-bot/internal/entity"

	"github.com/jackc/pgx/v4"
)

type ShoppingManager struct {
	Conn *pgx.Conn
}

func NewShoppingManager(conn *pgx.Conn) *ShoppingManager {
	return &ShoppingManager{Conn: conn}
}

func (sm *ShoppingManager) AddPurchase(purchase *entity.Purchase) error {
	return nil
}

func (sm *ShoppingManager) UpdatePurchase(purchase *entity.Purchase) error {
	return nil
}

func (sm *ShoppingManager) DeletePurchase(id uint) error {
	return nil
}

func (sm *ShoppingManager) GetPurchases() []entity.Purchase {
	return nil
}
