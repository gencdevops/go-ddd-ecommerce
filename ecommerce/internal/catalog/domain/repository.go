package domain

import "context"

type CatalogRepository interface {
	// Product operations
	SaveProduct(ctx context.Context, product *Product) error
	FindProductByID(ctx context.Context, id ProductID) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id ProductID) error

	// Category operations
	SaveCategory(ctx context.Context, category *Category) error
	FindCategoryByID(ctx context.Context, id CategoryID) (*Category, error)
	UpdateCategory(ctx context.Context, category *Category) error
	DeleteCategory(ctx context.Context, id CategoryID) error
}
