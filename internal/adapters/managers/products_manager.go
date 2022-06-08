package managers

import (
	"context"
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/jackc/pgx/v4"
)

type ProductsManager struct {
	conn *pgx.Conn
}

func NewProductsManager(conn *pgx.Conn) *ProductsManager {
	return &ProductsManager{conn: conn}
}

func (pm *ProductsManager) Add(ctx context.Context, product entity.Object) error {
	query := `INSERT INTO products (name, total_count, owner_id) VALUES ($1, $2, $3);`
	_, err := pm.conn.Exec(ctx, query, product.GetName(), product.GetCount(), product.GetOwnerId())
	return err
}

func (pm *ProductsManager) Update(ctx context.Context, newProduct entity.Object) error {
	query := `UPDATE products SET name=$1, total_count=$2 WHERE id=$3;`
	_, err := pm.conn.Exec(ctx, query, newProduct.GetName(), newProduct.GetCount(), newProduct.GetId())
	return err
}

func (pm *ProductsManager) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM products WHERE id=$1;`
	_, err := pm.conn.Exec(ctx, query, id)
	return err
}

func (pm *ProductsManager) Get(ctx context.Context, id uint, ownerId uint, groups []int32) (entity.Object, error) {
	query := `SELECT * FROM products WHERE id=$1 AND owner_id=$2`
	query += getGroupsQuery(groups, 3)

	args := []interface{}{id, ownerId}
	for _, val := range groups {
		args = append(args, val)
	}

	product := &entity.Product{}

	row := pm.conn.QueryRow(ctx, query, args...)
	err := row.Scan(&product.Id, &product.Name, &product.TotalCount, &product.OwnerId, &product.Groups)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pm *ProductsManager) GetAll(ctx context.Context, ownerId uint, groups []int32) ([]entity.Object, error) {
	query := `SELECT * FROM products WHERE owner_id = $1`
	query += getGroupsQuery(groups, 2)
	var products []entity.Object

	args := []interface{}{ownerId}
	for _, val := range groups {
		args = append(args, val)
	}

	rows, err := pm.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	products, err = toProductsList(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}
