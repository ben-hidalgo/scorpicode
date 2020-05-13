package soxierabbit

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"backend/pkg/rabbit"
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
)

// used to store the Repo in Context
type key int

// Key is the key for the repo in context; public for mock injection
const Key key = 0

// Channels .
type Channels struct {
	HatCreatedChannel   chan hatdao.Hat
	OrderCreatedChannel chan orderdao.Order
}

// Listen .
func Listen(jc *jazz.Connection, c Channels) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, Key, c)

	listener := &rabbit.Listener{
		Context: ctx,
	}

	go jc.ProcessQueue(rabbit.SoxieHatCreatedQ.Name(), listener.Wrap(ProcessHatsHatCreated))
}

// TODO: need to implement fanout to non-durable queues for multi-pod deployment

// ProcessHatsHatCreated .
func ProcessHatsHatCreated(ctx context.Context, msg []byte) error {
	logrus.Infof("soxierabbit.ProcessHatsHatCreated() msg=%s", string(msg))
	var h hatdao.Hat
	err := json.Unmarshal(msg, &h)
	if err != nil {
		return err
	}

	From(ctx).HatCreatedChannel <- h
	return nil
}

// From gets the web socket channel from the context
func From(ctx context.Context) Channels {
	switch v := ctx.Value(Key).(type) {
	case Channels:
		return v
	default:
		panic("soxierabbit.From() no value found")
	}
}
