package infrastructure

import (
    "context"
    "ddd-kata-golang/ecommerce/internal/payment/domain"
    "errors"
    "sync"
)

type InMemoryPaymentRepository struct {
    payments map[string]*domain.Payment
    mu       sync.RWMutex
}

func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
    return &InMemoryPaymentRepository{
        payments: make(map[string]*domain.Payment),
    }
}

func (r *InMemoryPaymentRepository) Save(ctx context.Context, payment *domain.Payment) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.payments[payment.ID().String()] = payment
    return nil
}

func (r *InMemoryPaymentRepository) FindByID(ctx context.Context, id domain.PaymentID) (*domain.Payment, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    if payment, exists := r.payments[id.String()]; exists {
        return payment, nil
    }
    return nil, errors.New("ödeme bulunamadı")
}

func (r *InMemoryPaymentRepository) Update(ctx context.Context, payment *domain.Payment) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.payments[payment.ID().String()] = payment
    return nil
}

func (r *InMemoryPaymentRepository) FindByOrderID(ctx context.Context, orderID string) (*domain.Payment, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    for _, payment := range r.payments {
        if payment.OrderID() == orderID {
            return payment, nil
        }
    }
    return nil, errors.New("ödeme bulunamadı")
}