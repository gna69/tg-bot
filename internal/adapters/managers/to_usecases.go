package managers

import (
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/jackc/pgx/v4"
)

func toPurchasesList(rows pgx.Rows) ([]entity.Object, error) {
	var err error
	var purchases []entity.Object

	for rows.Next() {
		var purchase entity.Purchase
		err = rows.Scan(
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

		purchases = append(purchases, &purchase)
	}

	return purchases, nil
}

func toProductsList(rows pgx.Rows) ([]entity.Object, error) {
	var err error
	var products []entity.Object

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.Id, &product.Name, &product.TotalCount)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}

func toRecipesList(rows pgx.Rows) ([]entity.Object, error) {
	var err error
	var recipes []entity.Object

	for rows.Next() {
		var recipe entity.Recipe
		err = rows.Scan(
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

		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}
