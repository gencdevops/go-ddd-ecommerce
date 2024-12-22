package domain

import (
	"ddd-kata-golang/ecommerce/pkg/errors"

	"github.com/google/uuid"
)

type PaymentID struct {
	value string
}

func NewPaymentID() PaymentID {
	return PaymentID{value: uuid.New().String()}
}

func (p PaymentID) String() string {
	return p.value
}

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
)

type Payment struct {
	id      PaymentID
	orderID string
	amount  float64
	status  PaymentStatus
}

func NewPayment(orderID string, amount float64) (*Payment, error) {
	if orderID == "" {
		return nil, errors.ErrOrderNotFound
	}
	if amount <= 0 {
		return nil, errors.ErrInvalidAmount
	}

	return &Payment{
		id:      NewPaymentID(),
		orderID: orderID,
		amount:  amount,
		status:  PaymentStatusPending,
	}, nil
}

func (p *Payment) ID() PaymentID {
	return p.id
}

func (p *Payment) OrderID() string {
	return p.orderID
}

func (p *Payment) Amount() float64 {
	return p.amount
}

func (p *Payment) Status() PaymentStatus {
	return p.status
}
