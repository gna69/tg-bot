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

func (rm *RecipesManager) Add(ctx context.Context, recipe entity.Object) error {
	query := `INSERT INTO recipes ("name", description, ingredients, actions, complexity, owner_id) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := rm.conn.Exec(ctx, query,
		recipe.GetName(),
		recipe.GetDescription(),
		recipe.GetIngredients(),
		recipe.GetActions(),
		recipe.GetComplexity(),
		recipe.GetOwnerId(),
	)
	return err
}

func (rm *RecipesManager) Update(ctx context.Context, newRecipe entity.Object) error {
	query := `UPDATE recipes SET "name"=$1, description=$2, ingredients=$3, actions=$4, complexity=$5 WHERE id=$6;`
	_, err := rm.conn.Exec(ctx, query,
		newRecipe.GetName(),
		newRecipe.GetDescription(),
		newRecipe.GetIngredients(),
		newRecipe.GetActions(),
		newRecipe.GetComplexity(),
		newRecipe.GetId(),
	)
	return err
}

func (rm *RecipesManager) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM recipes WHERE id=$1;`
	_, err := rm.conn.Exec(ctx, query, id)
	return err
}

func (rm *RecipesManager) Get(ctx context.Context, id uint, ownerId uint, groups []int32) (entity.Object, error) {
	query := `SELECT * FROM recipes WHERE id=$1 AND owner_id=$2`
	query += getGroupsQuery(groups, 3)

	args := []interface{}{id, ownerId}
	for _, val := range groups {
		args = append(args, val)
	}

	recipe := &entity.Recipe{}

	row := rm.conn.QueryRow(ctx, query, args...)
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

func (rm *RecipesManager) GetAll(ctx context.Context, ownerId uint, groups []int32) ([]entity.Object, error) {
	query := `SELECT * FROM recipes WHERE owner_id = $1`
	query += getGroupsQuery(groups, 2)

	args := []interface{}{ownerId}
	for _, val := range groups {
		args = append(args, val)
	}

	var recipes []entity.Object

	rows, err := rm.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	recipes, err = toRecipesList(rows)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
