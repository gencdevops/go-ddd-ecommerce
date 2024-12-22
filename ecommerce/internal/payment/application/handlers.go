package application

import (
    "encoding/json"
    "net/http"
)

type ProcessPaymentRequest struct {
    OrderID string  `json:"order_id"`
    Amount  float64 `json:"amount"`
}

type PaymentHandler struct {
    service *PaymentService
}

func NewPaymentHandler(service *PaymentService) *PaymentHandler {
    return &PaymentHandler{service: service}
}

func (h *PaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
    var req ProcessPaymentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    payment, err := h.service.ProcessPayment(r.Context(), req.OrderID, req.Amount)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(payment)
}