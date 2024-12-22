package main

import (
	"context"
	"ddd-kata-golang/ecommerce/internal/catalog/application"
	catalogInfra "ddd-kata-golang/ecommerce/internal/catalog/infrastructure"
	orderApp "ddd-kata-golang/ecommerce/internal/order/application"
	orderDomain "ddd-kata-golang/ecommerce/internal/order/domain"
	orderInfra "ddd-kata-golang/ecommerce/internal/order/infrastructure"
	paymentApp "ddd-kata-golang/ecommerce/internal/payment/application"
	paymentDomain "ddd-kata-golang/ecommerce/internal/payment/domain"
	paymentInfra "ddd-kata-golang/ecommerce/internal/payment/infrastructure"
	"ddd-kata-golang/ecommerce/internal/shared/domain"
	"ddd-kata-golang/ecommerce/internal/shared/infrastructure"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func setupEventHandlers(eventBus infrastructure.EventBus, orderService *orderApp.OrderService, paymentService *paymentApp.PaymentService) {
	eventBus.Subscribe("order.created", func(event domain.Event) {
		if orderEvent, ok := event.(*orderDomain.OrderCreatedEvent); ok {
			_, err := paymentService.ProcessPayment(context.Background(), orderEvent.OrderID, orderEvent.Total)
			if err != nil {
				log.Printf("Ödeme işlemi başarısız: %v", err)
			}
		}
	})

	eventBus.Subscribe("payment.processed", func(event domain.Event) {
		if paymentEvent, ok := event.(*paymentDomain.PaymentProcessedEvent); ok {
			log.Printf("Sipariş durumu güncellendi: %s", paymentEvent.OrderID)
		}
	})
}

func main() {

	eventBus := infrastructure.NewInMemoryEventBus()

	catalogRepo := catalogInfra.NewInMemoryCatalogRepository()
	orderRepo := orderInfra.NewInMemoryOrderRepository()
	paymentRepo := paymentInfra.NewInMemoryPaymentRepository()

	newCatalogService := application.NewCatalogService(catalogRepo, eventBus)
	orderService := orderApp.NewOrderService(orderRepo, eventBus)
	paymentService := paymentApp.NewPaymentService(paymentRepo, eventBus)

	productHandler := application.NewProductHandler(newCatalogService)
	orderHandler := orderApp.NewOrderHandler(orderService)
	paymentHandler := paymentApp.NewPaymentHandler(paymentService)

	setupEventHandlers(eventBus, orderService, paymentService)

	r := mux.NewRouter()
	r.Use(jsonMiddleware)
	r.Use(loggingMiddleware)

	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/payments", paymentHandler.ProcessPayment).Methods("POST")

	log.Printf("Server başlatılıyor: http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
