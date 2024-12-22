package domain_test

import (
	"ddd-kata-golang/ecommerce/internal/payment/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayment_Process(t *testing.T) {
	tests := []struct {
		name    string
		orderID string
		amount  float64
		wantErr bool
	}{
		{
			name:    "valid payment",
			orderID: "order-1",
			amount:  100.00,
			wantErr: false,
		},
		{
			name:    "invalid amount",
			orderID: "order-1",
			amount:  -100.00,
			wantErr: true,
		},
		{
			name:    "empty order id",
			orderID: "",
			amount:  100.00,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payment, err := domain.NewPayment(tt.orderID, tt.amount)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, payment)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, payment)
				assert.Equal(t, tt.orderID, payment.OrderID())
				assert.Equal(t, tt.amount, payment.Amount())
				assert.Equal(t, domain.PaymentStatusPending, payment.Status())
			}
		})
	}
}
