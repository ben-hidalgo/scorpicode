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
func (i *impl) Send(ex Exchange, key RKey, msg string) error {
	return i.Conn.SendMessage(string(ex), string(key), msg)
}

// SendBlob .
func (i *impl) SendBlob(ex Exchange, key RKey, msg []byte) error {
	return i.Conn.SendBlob(string(ex), string(key), msg)
}

// SendJSON .
func (i *impl) SendJSON(ex Exchange, key RKey, msg interface{}) error {

	s, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return i.SendBlob(ex, key, s)
}
