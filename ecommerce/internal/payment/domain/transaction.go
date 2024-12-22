package domain

import (
	"github.com/google/uuid"
	"time"
)

type TransactionID struct {
	value string
}

func NewTransactionID() TransactionID {
	return TransactionID{value: uuid.New().String()}
}

type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "PENDING"
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailed  TransactionStatus = "FAILED"
)

type Transaction struct {
	id        TransactionID
	amount    float64
	status    TransactionStatus
	createdAt time.Time
}
