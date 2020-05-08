package hatrabbit

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"backend/pkg/rabbit"
	"context"

	"github.com/socifi/jazz"
	"go.mongodb.org/mongo-driver/mongo"
)

// Listen .
func Listen(jc *jazz.Connection, mc *mongo.Client) {

	ctx := context.Background()

	// inject Rabbit Connection
	ctx = context.WithValue(ctx, rabbit.Key, rabbit.New(jc))

	// inject DAOs
	ctx = context.WithValue(ctx, hatdao.Key, hatdao.New(mc))
	ctx = context.WithValue(ctx, orderdao.Key, orderdao.New(mc))

	listener := &rabbit.Listener{
		Rabbit:  jc,
		Mongo:   mc,
		Context: ctx,
	}

	go jc.ProcessQueue(rabbit.HatsOrderCreatedQ.Name(), listener.Wrap(ProcessHatsOrderCreated))

	go jc.ProcessQueue(rabbit.HatsHatCreatedQ.Name(), listener.Wrap(ProcessHatsHatCreated))

}
