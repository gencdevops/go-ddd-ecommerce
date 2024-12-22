package domain

import (
	"ddd-kata-golang/ecommerce/internal/shared/domain"
	"time"
)

type OrderCreatedEvent struct {
	baseEvent  domain.BaseEvent
	OrderID    string  `json:"order_id"`
	CustomerID string  `json:"customer_id"`
	Total      float64 `json:"total"`
}

func NewOrderCreatedEvent(order *Order) *OrderCreatedEvent {
	return &OrderCreatedEvent{
		baseEvent:  domain.NewBaseEvent("order.created"),
		OrderID:    order.ID().String(),
		CustomerID: order.CustomerID(),
		Total:      order.Total(),
	}
}

// Event interface'ini implement etmek için gerekli metodlar
func (e *OrderCreatedEvent) Name() string {
	return e.baseEvent.Name()
}

func (e *OrderCreatedEvent) Timestamp() time.Time {
	return e.baseEvent.Timestamp()
}
