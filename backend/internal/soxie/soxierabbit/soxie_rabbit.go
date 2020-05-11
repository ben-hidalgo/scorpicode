package soxierabbit

import (
	"backend/pkg/rabbit"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
)

// used to store the Repo in Context
type key int

// Key is the key for the repo in context; public for mock injection
const Key key = 0

// Listen TODO: need to pass a handle to the socket writer wrapper
func Listen(jc *jazz.Connection, wsc chan string) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, Key, wsc)

	listener := &rabbit.Listener{
		Context: ctx,
	}

	go jc.ProcessQueue(rabbit.SoxieHatCreatedQ.Name(), listener.Wrap(ProcessHatsHatCreated))

}

// ProcessHatsHatCreated .
func ProcessHatsHatCreated(ctx context.Context, msg []byte) error {
	logrus.Infof("soxierabbit.ProcessHatsHatCreated() msg=%s", string(msg))
	wsc := From(ctx)
	tempString = fmt.Sprintf("%s\n%s", msg, tempString)
	wsc <- tempString
	return nil
}

var tempString = ""

// From gets the web socket channel from the context
func From(ctx context.Context) chan string {

	switch v := ctx.Value(Key).(type) {
	case chan string:
		return v
	default:
		panic("soxierabbit.From() no value found")
	}
}
