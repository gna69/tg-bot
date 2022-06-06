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

func (sm *ShoppingManager) Add(ctx context.Context, purchase entity.Object) error {
	query := `INSERT INTO purchases (name, description, count, unit, price, created_at) VALUES ($1, $2, $3, $4, $5, to_timestamp($6));`
	_, err := sm.conn.Exec(
		ctx,
		query,
		purchase.GetName(),
		purchase.GetDescription(),
		purchase.GetCount(),
		purchase.GetUnit(),
		purchase.GetPrice(),
		purchase.GetCreatedAt(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (sm *ShoppingManager) Update(ctx context.Context, newPurchase entity.Object) error {
	query := `UPDATE purchases SET name=$1, description=$2, count=$3, unit=$4, price=$5 WHERE id=$6;`
	_, err := sm.conn.Exec(
		ctx,
		query,
		newPurchase.GetName(),
		newPurchase.GetDescription(),
		newPurchase.GetCount(),
		newPurchase.GetUnit(),
		newPurchase.GetPrice(),
		newPurchase.GetId(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (sm *ShoppingManager) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM purchases WHERE id=$1;`
	_, err := sm.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (sm *ShoppingManager) Get(ctx context.Context, id uint) (entity.Object, error) {
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

func (sm *ShoppingManager) GetAll(ctx context.Context) ([]entity.Object, error) {
	query := `SELECT * FROM purchases;`
	var purchases []entity.Object

	rows, err := sm.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	purchases, err = toPurchasesList(rows)
	if err != nil {
		return nil, err
	}
	return purchases, nil
}
