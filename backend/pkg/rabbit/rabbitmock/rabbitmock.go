package rabbitmock

import (
	"backend/pkg/rabbit"
)

// Mock implements HatDao with all methods injectable
type Mock struct {
	SendF func(ex, msg, key string) error
}

// enforces the interface is implemented
var _ rabbit.Rmq = (*Mock)(nil)

// New .
func New() *Mock {
	return &Mock{}
}

// Send .
func (m *Mock) Send(ex, key, msg string) error {
	return m.SendF(ex, key, msg)
}
