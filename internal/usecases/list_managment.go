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
	GetPurchases(ctx context.Context) ([]entity.Purchase, error)
}

type ProductsManagement interface {
	AddProduct(product *entity.Product) error
	UpdateProduct(newProduct *entity.Product) error
	DeleteProduct(id uint) error
	GetProducts() []entity.Product
}

type RecipesManagement interface {
	AddRecipe(recipe *entity.Recipe) error
	UpdateRecipe(newRecipe *entity.Recipe) error
	DeleteRecipe(id uint) error
	GetRecipes() []entity.Recipe
}
