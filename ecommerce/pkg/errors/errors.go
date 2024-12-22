package errors

import (
	"fmt"
)

type DomainError struct {
	Code    string
	Message string
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewDomainError(code string, message string) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrInvalidProductName  = NewDomainError("INVALID_PRODUCT_NAME", "geçersiz ürün adı")
	ErrInvalidCategoryName = NewDomainError("INVALID_CATEGORY_NAME", "geçersiz kategori adı")
	ErrInvalidPrice        = NewDomainError("INVALID_PRICE", "geçersiz fiyat")
	ErrInvalidQuantity     = NewDomainError("INVALID_QUANTITY", "geçersiz miktar")
	ErrInvalidAmount       = NewDomainError("INVALID_AMOUNT", "geçersiz tutar")
	ErrProductNotFound     = NewDomainError("PRODUCT_NOT_FOUND", "ürün bulunamadı")
	ErrOrderNotFound       = NewDomainError("ORDER_NOT_FOUND", "sipariş bulunamadı")
	ErrPaymentNotFound     = NewDomainError("PAYMENT_NOT_FOUND", "ödeme bulunamadı")
	ErrInvalidCustomerID   = NewDomainError("INVALID_CUSTOMER_ID", "geçersiz müşteri ID")
	ErrInvalidProductID    = NewDomainError("INVALID_PRODUCT_ID", "geçersiz ürün ID")
	ErrInvalidOrderStatus  = NewDomainError("INVALID_ORDER_STATUS", "geçersiz sipariş durumu")
)
