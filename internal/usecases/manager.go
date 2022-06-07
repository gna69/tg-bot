package usecases

import (
	"context"

	"github.com/gna69/tg-bot/internal/entity"
)

var (
	EmptyList = "Этот список пустой, сначала добавь в него что-нибудь!"
)

type Manager interface {
	Add(ctx context.Context, obj entity.Object) error
	Update(ctx context.Context, newObj entity.Object) error
	Delete(ctx context.Context, objId uint) error
	Get(ctx context.Context, objId uint, ownerId uint) (entity.Object, error)
	GetAll(ctx context.Context, ownerId uint) ([]entity.Object, error)
}
