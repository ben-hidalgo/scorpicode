package hatrabbit

import (
	"backend/pkg/rabbit"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
)

// Listen .
func Listen(jc *jazz.Connection) {

	go jc.ProcessQueue(rabbit.HatsOrderCreatedQ.Name(), WrapProcessor(ProcessHatsOrderCreated))

}

// WrapProcessor wrappers the handler func with error handling
func WrapProcessor(f func(msg []byte) error) func(msg []byte) {

	handle := func(msg []byte) {

		err := f(msg)
		if err != nil {
			logrus.Errorf("hatrabbit.WrapProcessor() err=%#v", err)
			// TODO: publish to dead letter exchange
		}
	}
	return handle
}

// ProcessHatsOrderCreated handles
func ProcessHatsOrderCreated(msg []byte) error {
	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() msg=%s", string(msg))

	return nil
}
