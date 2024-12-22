package application

import (
	"context"
	"ddd-kata-golang/ecommerce/internal/catalog/domain"
	"ddd-kata-golang/ecommerce/internal/shared/infrastructure"
)

type CatalogService struct {
	repo     domain.CatalogRepository
	eventBus infrastructure.EventBus
}

func NewCatalogService(repo domain.CatalogRepository, eventBus infrastructure.EventBus) *CatalogService {
	return &CatalogService{
		repo:     repo,
		eventBus: eventBus,
	}
}

// Product operations
func (s *CatalogService) CreateProduct(ctx context.Context, name string, price float64, categoryID string) (*domain.Product, error) {
	product, err := domain.NewProduct(name, price, domain.CategoryID{Value: categoryID})
	if err != nil {
		return nil, err
	}

	if err := s.repo.SaveProduct(ctx, product); err != nil {
		return nil, err
	}

	event := domain.NewProductCreatedEvent(product)
	s.eventBus.Publish(event)

	return product, nil
}

func (s *CatalogService) GetProduct(ctx context.Context, id domain.ProductID) (*domain.Product, error) {
	return s.repo.FindProductByID(ctx, id)
}

func (s *CatalogService) UpdateProduct(ctx context.Context, product *domain.Product) error {
	return s.repo.UpdateProduct(ctx, product)
}

func (s *CatalogService) DeleteProduct(ctx context.Context, id domain.ProductID) error {
	return s.repo.DeleteProduct(ctx, id)
}

// Category operations
func (s *CatalogService) CreateCategory(ctx context.Context, name string, description string) (*domain.Category, error) {
	category, err := domain.NewCategory(name, description)
	if err != nil {
		return nil, err
	}

	if err := s.repo.SaveCategory(ctx, category); err != nil {
		return nil, err
	}

	event := domain.NewCategoryCreatedEvent(category)
	s.eventBus.Publish(event)

	return category, nil
}

func (s *CatalogService) GetCategory(ctx context.Context, id domain.CategoryID) (*domain.Category, error) {
	return s.repo.FindCategoryByID(ctx, id)
}

func (s *CatalogService) UpdateCategory(ctx context.Context, category *domain.Category) error {
	return s.repo.UpdateCategory(ctx, category)
}

func (s *CatalogService) DeleteCategory(ctx context.Context, id domain.CategoryID) error {
	return s.repo.DeleteCategory(ctx, id)
}
