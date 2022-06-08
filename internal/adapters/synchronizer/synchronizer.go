package synchronizer

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Synchronizer struct {
	db *pgx.Conn
}

func NewSynchronizer(db *pgx.Conn) *Synchronizer {
	return &Synchronizer{db: db}
}

func (s *Synchronizer) Synchronize(ctx context.Context, groups []int32, userId uint) error {
	shoppingQuery := `UPDATE purchases SET "groups" = $1 WHERE owner_id = $2;`
	productsQuery := `UPDATE products SET "groups" = $1 WHERE owner_id = $2;`
	recipesQuery := `UPDATE recipes SET "groups" = $1 WHERE owner_id = $2;`

	_, err := s.db.Exec(ctx, shoppingQuery, groups, userId)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, productsQuery, groups, userId)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, recipesQuery, groups, userId)
	if err != nil {
		return err
	}

	return nil
}
