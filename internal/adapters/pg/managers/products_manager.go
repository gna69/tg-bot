package managers

import (
	"github.com/gna69/tg-bot/internal/entity"

	"github.com/jackc/pgx/v4"
)

type ProductsManager struct {
	Conn *pgx.Conn
}

func NewProductsManager(conn *pgx.Conn) *ProductsManager {
	return &ProductsManager{Conn: conn}
}

func (pm *ProductsManager) AddProduct(product *entity.Product) error {
	return nil
}

func (pm *ProductsManager) UpdateProduct(newProduct *entity.Product) error {
	return nil
}

func (pm *ProductsManager) DeleteProduct(id uint) error {
	return nil
}

func (pm *ProductsManager) GetProducts() []entity.Product {
	return nil
}
