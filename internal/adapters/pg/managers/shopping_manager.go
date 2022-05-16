package managers

import (
	"context"
	"github.com/gna69/tg-bot/internal/entity"

	"github.com/jackc/pgx/v4"
)

type ShoppingManager struct {
	conn *pgx.Conn
}

func NewShoppingManager(conn *pgx.Conn) *ShoppingManager {
	return &ShoppingManager{conn: conn}
}

func (sm *ShoppingManager) AddPurchase(ctx context.Context, purchase *entity.Purchase) error {
	query := `INSERT INTO purchases (name, description, count, unit, price, created_at) VALUES ($1, $2, $3, $4, $5, to_timestamp($6));`
	_, err := sm.conn.Exec(ctx, query, purchase.Name, purchase.Description, purchase.Count, purchase.Unit, purchase.Price, purchase.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (sm *ShoppingManager) UpdatePurchase(ctx context.Context, newPurchase *entity.Purchase) error {
	query := `UPDATE purchases SET name=$1, description=$2, count=$3, unit=$4, price=$5, created_at=to_timestamp($6) WHERE id=$7;`
	_, err := sm.conn.Exec(ctx, query, newPurchase.Name, newPurchase.Description, newPurchase.Count, newPurchase.Unit, newPurchase.Price, newPurchase.CreatedAt, newPurchase.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sm *ShoppingManager) DeletePurchase(ctx context.Context, id uint) error {
	query := `DELETE FROM purchases WHERE id=$1;`
	_, err := sm.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (sm *ShoppingManager) GetPurchases(ctx context.Context) []entity.Purchase {
	return nil
}
