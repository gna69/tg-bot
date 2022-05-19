package managers

import (
	"github.com/gna69/tg-bot/internal/entity"
	"github.com/jackc/pgx/v4"
	"time"
)

func toPurchasesList(rows pgx.Row) ([]entity.Purchase, error) {
	var id uint
	var count uint8
	var price uint64
	var name, description, unit string
	var createdAt time.Time

	var err error
	var purchases []entity.Purchase
	for {
		err = rows.Scan(&id, &name, &description, &count, &unit, &price, &createdAt)

		if err == pgx.ErrNoRows {
			break
		}

		if err != nil {
			return nil, err
		}

		purchases = append(purchases, entity.Purchase{
			Id:          id,
			Count:       count,
			Price:       price,
			Name:        name,
			Description: description,
			Unit:        unit,
			CreatedAt:   createdAt,
		})
	}

	return purchases, nil
}
