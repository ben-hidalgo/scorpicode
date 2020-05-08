package rabbit

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
	"go.mongodb.org/mongo-driver/mongo"
)

// Listener is a stateful helper to facilitate message processing
type Listener struct {
	Rabbit  *jazz.Connection
	Mongo   *mongo.Client
	Context context.Context
}

// Processor is a func type to facilitate wrappering
type Processor func(ctx context.Context, msg []byte) error

// Wrap wrappers the handler func with error handling
func (l *Listener) Wrap(processor Processor) func(msg []byte) {

	// handler matches the required interface for jazz.ProcessQueue
	handler := func(msg []byte) {

		err := processor(l.Context, msg)
		if err != nil {
			logrus.Errorf("hatrabbit.WrapProcessor() err=%#v", err)
			// TODO: publish to dead letter exchange
		}
	}
	return handler
}
