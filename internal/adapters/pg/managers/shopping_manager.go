package managers

import (
	"context"
	"fmt"

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
	_, err := sm.conn.Exec(ctx, query, purchase.Name, purchase.Description, purchase.Count, purchase.Unit, purchase.Price, purchase.CreatedAt.Unix())
	if err != nil {
		return err
	}
	return nil
}

func (sm *ShoppingManager) UpdatePurchase(ctx context.Context, newPurchase *entity.Purchase) error {
	query := `UPDATE purchases SET name=$1, description=$2, count=$3, unit=$4, price=$5 WHERE id=$6;`
	_, err := sm.conn.Exec(ctx, query, newPurchase.Name, newPurchase.Description, newPurchase.Count, newPurchase.Unit, newPurchase.Price, newPurchase.Id)
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

func (sm *ShoppingManager) GetPurchase(ctx context.Context, id uint) (*entity.Purchase, error) {
	query := `SELECT * FROM purchases WHERE id=$1;`
	purchase := &entity.Purchase{}

	row := sm.conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&purchase.Id,
		&purchase.Name,
		&purchase.Description,
		&purchase.Count,
		&purchase.Unit,
		&purchase.Price,
		&purchase.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (sm *ShoppingManager) GetPurchases(ctx context.Context) ([]*entity.Purchase, error) {
	query := `SELECT * FROM purchases;`
	rows, err := sm.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	purchases, err := toPurchasesList(rows)
	if err != nil {
		return nil, err
	}
	return purchases, nil
}

func (sm *ShoppingManager) String(purchases []*entity.Purchase) string {
	list := ""
	for _, purchase := range purchases {
		list += fmt.Sprintf("\n%d) %s %d %s %d руб.", purchase.Id, purchase.Name, purchase.Count, purchase.Unit, purchase.Price)
		if len(purchase.Description) != 0 {
			list += fmt.Sprintf("\n%s", purchase.Description)
		}
		list += fmt.Sprintf("\n%s\n", purchase.CreatedAt.Format("15:04 02:January:2006"))
	}
	return list
}
