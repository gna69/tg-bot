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
