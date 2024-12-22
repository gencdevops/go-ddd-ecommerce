package infrastructure

import (
	"context"
	"ddd-kata-golang/ecommerce/internal/order/domain"
	"ddd-kata-golang/ecommerce/pkg/errors"
	"sync"
)

type InMemoryOrderRepository struct {
	orders map[string]*domain.Order
	mu     sync.RWMutex
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (r *InMemoryOrderRepository) Save(ctx context.Context, order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID().String()] = order
	return nil
}

func (r *InMemoryOrderRepository) FindByID(ctx context.Context, id domain.OrderID) (*domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if order, exists := r.orders[id.String()]; exists {
		return order, nil
	}
	return nil, errors.ErrOrderNotFound
}

func (r *InMemoryOrderRepository) Update(ctx context.Context, order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID().String()] = order
	return nil
}

func (r *InMemoryOrderRepository) Delete(ctx context.Context, id domain.OrderID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.orders, id.String())
	return nil
}
