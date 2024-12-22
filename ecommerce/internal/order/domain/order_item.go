package domain

import "ddd-kata-golang/ecommerce/pkg/errors"

type OrderItem struct {
	productID string
	quantity  int
	price     float64
}

func NewOrderItem(productID string, quantity int, price float64) (*OrderItem, error) {
	if productID == "" {
		return nil, errors.ErrInvalidOrderStatus
	}
	if quantity <= 0 {
		return nil, errors.ErrInvalidQuantity
	}
	if price <= 0 {
		return nil, errors.ErrInvalidPrice
	}

	return &OrderItem{
		productID: productID,
		quantity:  quantity,
		price:     price,
	}, nil
}

func (i *OrderItem) ProductID() string {
	return i.productID
}

func (i *OrderItem) Quantity() int {
	return i.quantity
}

func (i *OrderItem) Price() float64 {
	return i.price
}

func (i *OrderItem) Total() float64 {
	return float64(i.quantity) * i.price
}
