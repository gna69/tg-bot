package pg

import (
	"github.com/gna69/tg-bot/internal/adapters/pg/managers"
	"github.com/jackc/pgx/v4"
)

type PostgresClient struct {
	ShoppingManager *managers.ShoppingManager
	ProductsManager *managers.ProductsManager
	RecipesManager  *managers.RecipesManager
	WorkoutManager  *managers.WorkoutManager
}

func New(conn *pgx.Conn) *PostgresClient {
	return &PostgresClient{
		ShoppingManager: managers.NewShoppingManager(conn),
		ProductsManager: managers.NewProductsManager(conn),
		RecipesManager:  managers.NewRecipesManager(conn),
		WorkoutManager:  managers.NewWorkoutManager(conn),
	}
}
