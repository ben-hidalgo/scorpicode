package hatrabbit

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"backend/pkg/rabbit"
	"context"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
	"go.mongodb.org/mongo-driver/mongo"
)

// Listen .
func Listen(jc *jazz.Connection, mc *mongo.Client) {

	go jc.ProcessQueue(rabbit.HatsOrderCreatedQ.Name(), WrapProcessor(jc, mc, ProcessHatsOrderCreated))

}

// Processor is a func type to facilitate wrappering
type Processor func(ctx context.Context, msg []byte) error

// WrapProcessor wrappers the handler func with error handling
func WrapProcessor(jc *jazz.Connection, mc *mongo.Client, processor Processor) func(msg []byte) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, rabbit.Key, rabbit.New(jc))
	ctx = context.WithValue(ctx, hatdao.Key, hatdao.New(mc))
	ctx = context.WithValue(ctx, orderdao.Key, orderdao.New(mc))

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

// ProcessHatsOrderCreated handles
func ProcessHatsOrderCreated(ctx context.Context, msg []byte) error {

	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() msg=%s", string(msg))

	hd := hatdao.From(ctx)
	od := orderdao.From(ctx)
	rmq := rabbit.From(ctx)

	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() hd=%#v", hd)
	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() od=%#v", od)
	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() rmq=%#v", rmq)

	return nil
}
