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

func (pm *ProductsManager) AddProduct(ctx context.Context, product *entity.Product) error {
	query := `INSERT INTO products (name, total_count) VALUES ($1, $2);`
	_, err := pm.conn.Exec(ctx, query, product.Name, product.TotalCount)
	return err
}

func (pm *ProductsManager) UpdateProduct(ctx context.Context, newProduct *entity.Product) error {
	query := `UPDATE products SET name=$1, total_count=$2 WHERE id=$3;`
	_, err := pm.conn.Exec(ctx, query, newProduct.Name, newProduct.TotalCount, newProduct.Id)
	return err
}

func (pm *ProductsManager) DeleteProduct(ctx context.Context, id uint) error {
	query := `DELETE FROM products WHERE id=$1;`
	_, err := pm.conn.Exec(ctx, query, id)
	return err
}

func (pm *ProductsManager) GetProduct(ctx context.Context, id uint) (*entity.Product, error) {
	query := `SELECT * FROM products WHERE id=$1;`
	product := &entity.Product{}

	row := pm.conn.QueryRow(ctx, query, id)
	err := row.Scan(&product.Id, &product.Name, &product.TotalCount)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pm *ProductsManager) GetProducts(ctx context.Context) ([]*entity.Product, error) {
	query := `SELECT * FROM products;`
	rows, err := pm.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	products, err := toProductsList(rows)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pm *ProductsManager) String(products []*entity.Product) string {
	list := ""
	for _, product := range products {
		list += product.ToString()
	}
	return list
}
