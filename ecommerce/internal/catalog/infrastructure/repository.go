package infrastructure

import (
	"context"
	"ddd-kata-golang/ecommerce/internal/catalog/domain"
	"errors"
	"sync"
)

type InMemoryCatalogRepository struct {
	products   map[string]*domain.Product
	categories map[string]*domain.Category
	mu         sync.RWMutex
}

func NewInMemoryCatalogRepository() *InMemoryCatalogRepository {
	return &InMemoryCatalogRepository{
		products:   make(map[string]*domain.Product),
		categories: make(map[string]*domain.Category),
	}
}

// Product operations
func (r *InMemoryCatalogRepository) SaveProduct(ctx context.Context, product *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products[product.ID().String()] = product
	return nil
}

func (r *InMemoryCatalogRepository) FindProductByID(ctx context.Context, id domain.ProductID) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if product, exists := r.products[id.String()]; exists {
		return product, nil
	}
	return nil, errors.New("ürün bulunamadı")
}

func (r *InMemoryCatalogRepository) UpdateProduct(ctx context.Context, product *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products[product.ID().String()] = product
	return nil
}

func (r *InMemoryCatalogRepository) DeleteProduct(ctx context.Context, id domain.ProductID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.products, id.String())
	return nil
}

// Category operations
func (r *InMemoryCatalogRepository) SaveCategory(ctx context.Context, category *domain.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.categories[category.ID().String()] = category
	return nil
}

func (r *InMemoryCatalogRepository) FindCategoryByID(ctx context.Context, id domain.CategoryID) (*domain.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if category, exists := r.categories[id.String()]; exists {
		return category, nil
	}
	return nil, errors.New("kategori bulunamadı")
}

func (r *InMemoryCatalogRepository) UpdateCategory(ctx context.Context, category *domain.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.categories[category.ID().String()] = category
	return nil
}

func (r *InMemoryCatalogRepository) DeleteCategory(ctx context.Context, id domain.CategoryID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.categories, id.String())
	return nil
}
