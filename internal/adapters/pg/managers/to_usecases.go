package managers

import (
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/jackc/pgx/v4"
)

func toPurchasesList(rows pgx.Rows) ([]*entity.Purchase, error) {
	var err error
	var purchases []*entity.Purchase

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

func toProductsList(rows pgx.Rows) ([]*entity.Product, error) {
	var err error
	var products []*entity.Product

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

func toRecipesList(rows pgx.Rows) ([]*entity.Recipe, error) {
	var err error
	var recipes []*entity.Recipe

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
