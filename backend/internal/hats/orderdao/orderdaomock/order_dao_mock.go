package orderdaomock

import (
	"backend/internal/hats/orderdao"
	"context"
)

// Mock implements OrderDao with all methods injectable
type Mock struct {
	CreateF func(context.Context, *orderdao.Order) error
	FindF   func(context.Context, string) (*orderdao.Order, error)
}

// enforces the interface is implemented
var _ orderdao.OrderDao = (*Mock)(nil)

// New .
func New() *Mock {
	return &Mock{}
}

// Create calls the injected function
func (r *Mock) Create(ctx context.Context, h *orderdao.Order) error {
	return r.CreateF(ctx, h)
}

// Find calls the injected function
func (r *Mock) Find(ctx context.Context, id string) (*orderdao.Order, error) {
	return r.FindF(ctx, id)
}
