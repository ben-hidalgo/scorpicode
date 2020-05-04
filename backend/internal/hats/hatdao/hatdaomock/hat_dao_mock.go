package hatdaomock

import (
	"backend/internal/hats/hatdao"
	"context"
)

// Mock implements HatDao with all methods injectable
type Mock struct {
	CreateF func(context.Context, *hatdao.Hat) error
	DeleteF func(context.Context, *hatdao.Hat) error
	FindF   func(context.Context, string) (*hatdao.Hat, error)
	QueryF  func(context.Context) ([]*hatdao.Hat, error)

	VisitTxnF func(context.Context, func() error) error
}

// enforces the interface is implemented
var _ hatdao.HatDao = (*Mock)(nil)

// New .
func New() *Mock {
	return &Mock{}
}

// Create calls the injected function
func (r *Mock) Create(ctx context.Context, h *hatdao.Hat) error {
	return r.CreateF(ctx, h)
}

// Delete calls the injected function
func (r *Mock) Delete(ctx context.Context, h *hatdao.Hat) error {
	return r.DeleteF(ctx, h)
}

// Find calls the injected function
func (r *Mock) Find(ctx context.Context, id string) (*hatdao.Hat, error) {
	return r.FindF(ctx, id)
}

// Query calls the injected function
func (r *Mock) Query(ctx context.Context) ([]*hatdao.Hat, error) {
	return r.QueryF(ctx)
}

// VisitTxn .
func (r *Mock) VisitTxn(ctx context.Context, tf func() error) error {
	return r.VisitTxnF(ctx, tf)
}
