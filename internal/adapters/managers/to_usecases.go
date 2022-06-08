package managers

import (
	"fmt"
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/jackc/pgx/v4"
)

var groups []int

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
			&purchase.OwnerId,
			&groups,
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
		err = rows.Scan(&product.Id, &product.Name, &product.TotalCount, &product.OwnerId, &groups)
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
			&recipe.OwnerId,
			&groups,
		)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}

func getGroupsQuery(groups []int32, argIdx int) (query string) {
	if len(groups) == 0 {
		return ";"
	}

	query += fmt.Sprintf(` OR $%d = ANY ("groups") `, argIdx)
	groups[0] = groups[len(groups)-1]
	groups = groups[:len(groups)-1]

	for range groups {
		argIdx++
		query += fmt.Sprintf(`OR $%d = ANY ("groups") `, argIdx)
	}
	query += ";"
	return
}
