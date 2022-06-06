package pg

import (
	"github.com/gna69/tg-bot/internal/adapters/managers"
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/gna69/tg-bot/internal/usecases"
	"github.com/jackc/pgx/v4"
)

func NewManager(botMode string, conn *pgx.Conn) usecases.Manager {
	var manager usecases.Manager

	switch botMode {
	case entity.Shopping:
		manager = managers.NewShoppingManager(conn)
	case entity.Products:
		manager = managers.NewProductsManager(conn)
	case entity.Recipes:
		manager = managers.NewRecipesManager(conn)
	}

	return manager
}
