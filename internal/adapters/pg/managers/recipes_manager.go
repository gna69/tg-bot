package managers

import (
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/jackc/pgx/v4"
)

type RecipesManager struct {
	Conn *pgx.Conn
}

func NewRecipesManager(conn *pgx.Conn) *RecipesManager {
	return &RecipesManager{Conn: conn}
}

func (r *RecipesManager) AddRecipe(recipe *entity.Recipe) error {
	return nil
}

func (r *RecipesManager) UpdateRecipe(newRecipe *entity.Recipe) error {
	return nil
}

func (r *RecipesManager) DeleteRecipe(id uint) error {
	return nil
}

func (r *RecipesManager) GetRecipes() []entity.Recipe {
	return nil
}
