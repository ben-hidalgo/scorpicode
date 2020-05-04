package rabbit

import (
	"encoding/json"

	"github.com/socifi/jazz"
)

// AmqpDsn .

//impl .
type impl struct {
	Conn *jazz.Connection
}

// Send .
func (i *impl) Send(ex, key, msg string) error {
	return i.Conn.SendMessage(ex, key, msg)
}

// SendBlob .
func (i *impl) SendBlob(ex, key string, msg []byte) error {
	return i.Conn.SendBlob(ex, key, msg)
}

// SendJSON .
func (i *impl) SendJSON(ex, key string, msg interface{}) error {

	s, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return i.SendBlob(ex, key, s)
}
