package application

import (
	"context"
	"ddd-kata-golang/ecommerce/internal/payment/domain"
	"ddd-kata-golang/ecommerce/internal/shared/infrastructure"
)

type PaymentService struct {
	repo     domain.PaymentRepository
	eventBus infrastructure.EventBus
}

func NewPaymentService(repo domain.PaymentRepository, eventBus infrastructure.EventBus) *PaymentService {
	return &PaymentService{
		repo:     repo,
		eventBus: eventBus,
	}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, orderID string, amount float64) (*domain.Payment, error) {
	payment, err := domain.NewPayment(orderID, amount)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, payment); err != nil {
		return nil, err
	}

	event := domain.NewPaymentProcessedEvent(payment)
	s.eventBus.Publish(event)

	return payment, nil
}
