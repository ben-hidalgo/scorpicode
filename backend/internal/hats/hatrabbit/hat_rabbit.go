package hatrabbit

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"backend/pkg/rabbit"
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
	"go.mongodb.org/mongo-driver/mongo"
)

// Listen .
func Listen(jc *jazz.Connection, mc *mongo.Client) {

	go jc.ProcessQueue(rabbit.HatsOrderCreatedQ.Name(), WrapProcessor(jc, mc, ProcessHatsOrderCreated))

	go jc.ProcessQueue(rabbit.HatsHatCreatedQ.Name(), WrapProcessor(jc, mc, ProcessHatsHatCreated))

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

// ProcessHatsOrderCreated .
func ProcessHatsOrderCreated(ctx context.Context, msg []byte) error {

	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() msg=%s", string(msg))

	hd := hatdao.From(ctx)
	od := orderdao.From(ctx)
	rmq := rabbit.From(ctx)

	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() hd=%#v", hd)
	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() od=%#v", od)
	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() rmq=%#v", rmq)

	var order orderdao.Order
	err := json.Unmarshal(msg, &order)
	if err != nil {
		return err
	}
	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() order=%#v", order)

	// this is the transaction func
	tf := func() error {

		// save a hat per quantity
		for i := int32(0); i < order.Quantity; i++ {
			hat := &hatdao.Hat{
				Color: order.Color,
				Style: order.Style,
				Size:  order.Size,
			}
			err := hd.Create(ctx, hat)
			if err != nil {
				return err
			}

			rmq.SendJSON(rabbit.ServiceMsgActionX, rabbit.HatsHatCreatedK, hat)
		}
		return nil
	}

	err = hd.VisitTxn(ctx, tf)
	if err != nil {
		return err
	}

	return nil
}

// ProcessHatsHatCreated .
func ProcessHatsHatCreated(ctx context.Context, msg []byte) error {
	logrus.Infof("hatrabbit.ProcessHatsHatCreated() msg=%s", string(msg))
	return nil
}
