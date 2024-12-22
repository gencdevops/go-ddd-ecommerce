package application

import (
	"context"
	"ddd-kata-golang/ecommerce/internal/order/domain"
	"ddd-kata-golang/ecommerce/internal/shared/infrastructure"
)

type OrderService struct {
	repository domain.OrderRepository
	eventBus   infrastructure.EventBus
}

func NewOrderService(repository domain.OrderRepository, eventBus infrastructure.EventBus) *OrderService {
	return &OrderService{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, customerID string, items []*domain.OrderItem) (*domain.Order, error) {
	order, err := domain.NewOrder(customerID, items)
	if err != nil {
		return nil, err
	}

	if err := s.repository.Save(ctx, order); err != nil {
		return nil, err
	}

	event := domain.NewOrderCreatedEvent(order)
	s.eventBus.Publish(event)

	return order, nil
}

func (s *OrderService) GetOrder(ctx context.Context, orderID domain.OrderID) (*domain.Order, error) {
	return s.repository.FindByID(ctx, orderID)
}
