package usecases

import (
	"context"
	"github.com/gna69/tg-bot/internal/entity"
)

var (
	EmptyList = "Этот список пустой, сначала добавь в него что-нибудь!"
)

type ShoppingManagement interface {
	AddPurchase(ctx context.Context, purchase *entity.Purchase) error
	UpdatePurchase(ctx context.Context, newPurchase *entity.Purchase) error
	DeletePurchase(ctx context.Context, id uint) error
	GetPurchase(ctx context.Context, id uint) (*entity.Purchase, error)
	GetPurchases(ctx context.Context) ([]*entity.Purchase, error)
}

type ProductsManagement interface {
	AddProduct(ctx context.Context, product *entity.Product) error
	UpdateProduct(ctx context.Context, newProduct *entity.Product) error
	DeleteProduct(ctx context.Context, id uint) error
	GetProduct(ctx context.Context, id uint) (*entity.Product, error)
	GetProducts(ctx context.Context) ([]*entity.Product, error)
}

type RecipesManagement interface {
	AddRecipe(ctx context.Context, recipe *entity.Recipe) error
	UpdateRecipe(ctx context.Context, newRecipe *entity.Recipe) error
	DeleteRecipe(ctx context.Context, id uint) error
	GetRecipe(ctx context.Context, id uint) (*entity.Recipe, error)
	GetRecipes(ctx context.Context) ([]*entity.Recipe, error)
}
