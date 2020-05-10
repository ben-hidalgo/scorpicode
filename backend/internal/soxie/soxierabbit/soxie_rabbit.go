package soxierabbit

import (
	"backend/pkg/rabbit"
	"context"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
)

// Listen TODO: need to pass a handle to the socket writer wrapper
func Listen(jc *jazz.Connection) {

	ctx := context.Background()

	listener := &rabbit.Listener{
		Context: ctx,
	}

	go jc.ProcessQueue(rabbit.SoxieHatCreatedQ.Name(), listener.Wrap(ProcessHatsHatCreated))

}

// ProcessHatsHatCreated .
func ProcessHatsHatCreated(ctx context.Context, msg []byte) error {
	logrus.Infof("soxierabbit.ProcessHatsHatCreated() msg=%s", string(msg))
	return nil
}
