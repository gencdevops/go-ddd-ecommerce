package domain

import (
	"ddd-kata-golang/ecommerce/pkg/errors"
	"github.com/google/uuid"
)

type OrderID struct {
	value string
}

func NewOrderID() OrderID {
	return OrderID{value: uuid.New().String()}
}

func (o OrderID) String() string {
	return o.value
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

var validStatusTransitions = map[OrderStatus][]OrderStatus{
	OrderStatusPending: {
		OrderStatusPaid,
		OrderStatusCancelled,
	},
	OrderStatusPaid: {
		OrderStatusShipped,
		OrderStatusCancelled,
	},
	OrderStatusShipped: {
		OrderStatusDelivered,
		OrderStatusCancelled,
	},
	OrderStatusDelivered: {
		OrderStatusCancelled,
	},
}

type Order struct {
	id         OrderID
	customerID string
	items      []*OrderItem
	total      float64
	status     OrderStatus
}

func NewOrder(customerID string, items []*OrderItem) (*Order, error) {
	if customerID == "" {
		return nil, errors.ErrInvalidCustomerID
	}

	if len(items) == 0 {
		return nil, errors.ErrInvalidQuantity
	}

	total := 0.0
	for _, item := range items {
		if item == nil {
			return nil, errors.ErrInvalidQuantity
		}
		if item.Quantity() <= 0 {
			return nil, errors.ErrInvalidQuantity
		}
		total += item.Total()
	}

	return &Order{
		id:         NewOrderID(),
		customerID: customerID,
		items:      items,
		total:      total,
		status:     OrderStatusPending,
	}, nil
}

func (o *Order) ID() OrderID {
	return o.id
}

func (o *Order) CustomerID() string {
	return o.customerID
}

func (o *Order) Items() []*OrderItem {
	return o.items
}

func (o *Order) Total() float64 {
	return o.total
}

func (o *Order) Status() OrderStatus {
	return o.status
}

func (o *Order) UpdateStatus(newStatus OrderStatus) error {
	if o.status == OrderStatusCancelled {
		return errors.ErrInvalidOrderStatus
	}

	validNextStatuses, exists := validStatusTransitions[o.status]
	if !exists {
		return errors.ErrInvalidOrderStatus
	}

	for _, validStatus := range validNextStatuses {
		if newStatus == validStatus {
			o.status = newStatus
			return nil
		}
	}

	return errors.ErrInvalidOrderStatus
}
