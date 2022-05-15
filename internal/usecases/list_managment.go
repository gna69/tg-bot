package usecases

import "github.com/gna69/tg-bot/internal/entity"

type ShoppingManagement interface {
	AddPurchase(purchase *entity.Purchase) error
	UpdatePurchase(newPurchase *entity.Purchase) error
	DeletePurchase(id uint) error
	GetPurchases() []entity.Purchase
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
