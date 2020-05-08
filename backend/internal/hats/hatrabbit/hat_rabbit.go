package hatrabbit

import (
	"backend/pkg/rabbit"

	"github.com/socifi/jazz"
	"go.mongodb.org/mongo-driver/mongo"
)

// Listen .
func Listen(jc *jazz.Connection, mc *mongo.Client) {

	listener := &rabbit.Listener{
		Rabbit: jc,
		Mongo:  mc,
	}

	go jc.ProcessQueue(rabbit.HatsOrderCreatedQ.Name(), listener.Wrap(ProcessHatsOrderCreated))

	go jc.ProcessQueue(rabbit.HatsHatCreatedQ.Name(), listener.Wrap(ProcessHatsHatCreated))

}
