package rabbit

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"context"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
	"go.mongodb.org/mongo-driver/mongo"
)

// Listener is a stateful helper to facilitate message processing
type Listener struct {
	Rabbit *jazz.Connection
	Mongo  *mongo.Client
}

// Processor is a func type to facilitate wrappering
type Processor func(ctx context.Context, msg []byte) error

// Wrap wrappers the handler func with error handling
func (l *Listener) Wrap(processor Processor) func(msg []byte) {

	ctx := context.Background()

	// inject Rabbit Connection
	ctx = context.WithValue(ctx, Key, New(l.Rabbit))

	// inject DAOs
	ctx = context.WithValue(ctx, hatdao.Key, hatdao.New(l.Mongo))
	ctx = context.WithValue(ctx, orderdao.Key, orderdao.New(l.Mongo))

	// TODO: inject DAOs and RMQ into context

	// handler matches the required interface for jazz.ProcessQueue
	handler := func(msg []byte) {

		err := processor(ctx, msg)
		if err != nil {
			logrus.Errorf("hatrabbit.WrapProcessor() err=%#v", err)
			// TODO: publish to dead letter exchange
		}
	}
	return handler
}
