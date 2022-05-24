package managers

import (
	"context"

	"github.com/gna69/tg-bot/internal/entity"
	"github.com/jackc/pgx/v4"
)

type RecipesManager struct {
	conn *pgx.Conn
}

func NewRecipesManager(conn *pgx.Conn) *RecipesManager {
	return &RecipesManager{conn: conn}
}

func (rm *RecipesManager) AddRecipe(ctx context.Context, recipe *entity.Recipe) error {
	query := `INSERT INTO recipes ("name", description, ingredients, actions, complexity) VALUES ($1, $2, $3, $4, $5);`
	_, err := rm.conn.Exec(ctx, query,
		recipe.Name,
		recipe.Description,
		recipe.Ingredients,
		recipe.Actions,
		recipe.Complexity,
	)
	return err
}

func (rm *RecipesManager) UpdateRecipe(ctx context.Context, newRecipe *entity.Recipe) error {
	query := `UPDATE recipes SET "name"=$1, description=$2, ingredients=$3, actions=$4, complexity=$5 WHERE id=$6;`
	_, err := rm.conn.Exec(ctx, query,
		newRecipe.Name,
		newRecipe.Description,
		newRecipe.Ingredients,
		newRecipe.Actions,
		newRecipe.Complexity,
		newRecipe.Id,
	)
	return err
}

func (rm *RecipesManager) DeleteRecipe(ctx context.Context, id uint) error {
	query := `DELETE FROM recipes WHERE id=$1;`
	_, err := rm.conn.Exec(ctx, query, id)
	return err
}

func (rm *RecipesManager) GetRecipe(ctx context.Context, id uint) (*entity.Recipe, error) {
	query := `SELECT * FROM recipes WHERE id=$1;`
	recipe := &entity.Recipe{}

	row := rm.conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&recipe.Id,
		&recipe.Name,
		&recipe.Description,
		&recipe.Ingredients,
		&recipe.Actions,
		&recipe.Complexity,
	)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (rm *RecipesManager) GetRecipes(ctx context.Context) ([]*entity.Recipe, error) {
	query := `SELECT * FROM recipes;`
	rows, err := rm.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	recipes, err := toRecipesList(rows)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (rm *RecipesManager) String(recipes []*entity.Recipe) string {
	list := ""
	for _, recipe := range recipes {
		list += recipe.ToString()
	}
	return list
}
