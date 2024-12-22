package domain

import (
	"ddd-kata-golang/ecommerce/internal/shared/domain"
	"time"
)

type PaymentProcessedEvent struct {
	baseEvent domain.BaseEvent
	PaymentID string        `json:"payment_id"`
	OrderID   string        `json:"order_id"`
	Amount    float64       `json:"amount"`
	Status    PaymentStatus `json:"status"`
}

func NewPaymentProcessedEvent(payment *Payment) *PaymentProcessedEvent {
	return &PaymentProcessedEvent{
		baseEvent: domain.NewBaseEvent("payment.processed"),
		PaymentID: payment.ID().String(),
		OrderID:   payment.OrderID(),
		Amount:    payment.Amount(),
		Status:    payment.Status(),
	}
}

// Event interface'ini implement etmek için gerekli metodlar
func (e *PaymentProcessedEvent) Name() string {
	return e.baseEvent.Name()
}

func (e *PaymentProcessedEvent) Timestamp() time.Time {
	return e.baseEvent.Timestamp()
}
