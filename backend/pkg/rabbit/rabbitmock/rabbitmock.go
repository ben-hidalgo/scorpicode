package rabbitmock

import (
	"backend/pkg/rabbit"
)

// Mock implements HatDao with all methods injectable
type Mock struct {
	SendF     func(ex rabbit.Exchange, key rabbit.RKey, msg string) error
	SendBlobF func(ex rabbit.Exchange, key rabbit.RKey, msg []byte) error
	SendJSONF func(ex rabbit.Exchange, key rabbit.RKey, msg interface{}) error
}

// enforces the interface is implemented
var _ rabbit.Rmq = (*Mock)(nil)

// New .
func New() *Mock {
	return &Mock{}
}

// Send .
func (m *Mock) Send(ex rabbit.Exchange, key rabbit.RKey, msg string) error {
	return m.SendF(ex, key, msg)
}

// SendBlob .
func (m *Mock) SendBlob(ex rabbit.Exchange, key rabbit.RKey, msg []byte) error {
	return m.SendBlobF(ex, key, msg)
}

// SendJSON .
func (m *Mock) SendJSON(ex rabbit.Exchange, key rabbit.RKey, msg interface{}) error {
	return m.SendJSONF(ex, key, msg)
}
