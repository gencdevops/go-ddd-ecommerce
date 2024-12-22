package domain

import (
	"ddd-kata-golang/ecommerce/internal/shared/domain"
	"time"
)

type ProductCreatedEvent struct {
	baseEvent   domain.BaseEvent
	ProductID   string  `json:"product_id"`
	ProductName string  `json:"name"`
	Price       float64 `json:"price"`
}

func NewProductCreatedEvent(product *Product) *ProductCreatedEvent {
	return &ProductCreatedEvent{
		baseEvent:   domain.NewBaseEvent("product.created"),
		ProductID:   product.ID().String(),
		ProductName: product.Name(),
		Price:       product.Price(),
	}
}

// Event interface'ini implement etmek için gerekli metodlar
func (e *ProductCreatedEvent) Name() string {
	return e.baseEvent.Name()
}

func (e *ProductCreatedEvent) Timestamp() time.Time {
	return e.baseEvent.Timestamp()
}
