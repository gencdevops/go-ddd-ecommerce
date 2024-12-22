package infrastructure

import (
	"ddd-kata-golang/ecommerce/internal/shared/domain"
	"sync"
)

type EventBus interface {
	Publish(event domain.Event)
	Subscribe(eventName string, handler func(event domain.Event))
}

type InMemoryEventBus struct {
	handlers map[string][]func(event domain.Event)
	mu       sync.RWMutex
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		handlers: make(map[string][]func(event domain.Event)),
	}
}

func (b *InMemoryEventBus) Publish(event domain.Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if handlers, exists := b.handlers[event.Name()]; exists {
		for _, handler := range handlers {
			go handler(event)
		}
	}
}

func (b *InMemoryEventBus) Subscribe(eventName string, handler func(event domain.Event)) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[eventName] = append(b.handlers[eventName], handler)
}
