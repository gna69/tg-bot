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
	query := `INSERT INTO products (name, total_count) VALUES ($1, $2);`
	_, err := pm.conn.Exec(ctx, query, product.GetName(), product.GetCount())
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

func (pm *ProductsManager) Get(ctx context.Context, id uint) (entity.Object, error) {
	query := `SELECT * FROM products WHERE id=$1;`
	product := &entity.Product{}

	row := pm.conn.QueryRow(ctx, query, id)
	err := row.Scan(&product.Id, &product.Name, &product.TotalCount)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pm *ProductsManager) GetAll(ctx context.Context) ([]entity.Object, error) {
	query := `SELECT * FROM products;`
	var products []entity.Object

	rows, err := pm.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	products, err = toProductsList(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}
