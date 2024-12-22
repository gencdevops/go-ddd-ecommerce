package domain_test

import (
	"ddd-kata-golang/ecommerce/internal/order/domain"
	"ddd-kata-golang/ecommerce/pkg/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder_Create(t *testing.T) {
	validItem, _ := domain.NewOrderItem("product-1", 2, 100.00)

	tests := []struct {
		name       string
		customerID string
		items      []*domain.OrderItem
		wantErr    error
	}{
		{
			name:       "valid order",
			customerID: "customer-1",
			items:      []*domain.OrderItem{validItem},
			wantErr:    nil,
		},
		{
			name:       "empty customer id",
			customerID: "",
			items:      []*domain.OrderItem{validItem},
			wantErr:    errors.ErrInvalidCustomerID,
		},
		{
			name:       "empty items",
			customerID: "customer-1",
			items:      []*domain.OrderItem{},
			wantErr:    errors.ErrInvalidQuantity,
		},
		{
			name:       "nil items",
			customerID: "customer-1",
			items:      []*domain.OrderItem{nil},
			wantErr:    errors.ErrInvalidQuantity,
		},
		{
			name:       "multiple items",
			customerID: "customer-1",
			items:      []*domain.OrderItem{validItem, validItem},
			wantErr:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, err := domain.NewOrder(tt.customerID, tt.items)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr, err)
				assert.Nil(t, order)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, order)
				assert.Equal(t, tt.customerID, order.CustomerID())
				assert.Equal(t, tt.items, order.Items())
				assert.Equal(t, domain.OrderStatusPending, order.Status())

				expectedTotal := 0.0
				for _, item := range tt.items {
					expectedTotal += item.Total()
				}
				assert.Equal(t, expectedTotal, order.Total())
			}
		})
	}
}

func TestOrder_UpdateStatus(t *testing.T) {
	validItem, _ := domain.NewOrderItem("product-1", 2, 100.00)
	order, _ := domain.NewOrder("customer-1", []*domain.OrderItem{validItem})

	tests := []struct {
		name      string
		oldStatus domain.OrderStatus
		newStatus domain.OrderStatus
		wantErr   bool
	}{
		{
			name:      "pending to paid",
			oldStatus: domain.OrderStatusPending,
			newStatus: domain.OrderStatusPaid,
			wantErr:   false,
		},
		{
			name:      "pending to cancelled",
			oldStatus: domain.OrderStatusPending,
			newStatus: domain.OrderStatusCancelled,
			wantErr:   false,
		},
		{
			name:      "pending to shipped",
			oldStatus: domain.OrderStatusPending,
			newStatus: domain.OrderStatusShipped,
			wantErr:   true,
		},
		{
			name:      "cancelled to paid",
			oldStatus: domain.OrderStatusCancelled,
			newStatus: domain.OrderStatusPaid,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order.UpdateStatus(tt.oldStatus) // Set initial status
			err := order.UpdateStatus(tt.newStatus)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, errors.ErrInvalidOrderStatus, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newStatus, order.Status())
			}
		})
	}
}
