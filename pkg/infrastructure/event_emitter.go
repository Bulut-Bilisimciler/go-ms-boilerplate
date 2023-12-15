package infrastructure

import (
	"context"
	"sync"
)

type EventEmitter struct {
	mu       sync.Mutex
	handlers map[string][]func(context.Context, interface{})
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		handlers: make(map[string][]func(context.Context, interface{})),
	}
}

func (e *EventEmitter) On(eventName string, handler func(context.Context, interface{})) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, ok := e.handlers[eventName]; !ok {
		e.handlers[eventName] = make([]func(context.Context, interface{}), 0)
	}

	e.handlers[eventName] = append(e.handlers[eventName], handler)
}

// EmitAsync emits an event asynchronously
func (e *EventEmitter) EmitAsync(ctx context.Context, eventName string, data interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if handlers, ok := e.handlers[eventName]; ok {
		for _, handler := range handlers {
			go handler(ctx, data) // Asenkron olarak event'i işle
		}
	}
}

// Emit emits an event synchronously
func (e *EventEmitter) EmitSync(ctx context.Context, eventName string, data interface{}) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if handlers, ok := e.handlers[eventName]; ok {
		for _, handler := range handlers {
			handler(ctx, data) // Senkron olarak event'i işle
		}
	}
}
