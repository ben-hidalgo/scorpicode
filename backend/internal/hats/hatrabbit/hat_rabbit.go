package hatrabbit

import (
	"backend/pkg/rabbit"

	"github.com/sirupsen/logrus"
	"github.com/socifi/jazz"
)

// Listen .
func Listen(jc *jazz.Connection) {

	// Handler function
	f := func(msg []byte) {
		// TODO: add channels and switch by message type
		logrus.Infof("hatrabbit.Connect() msg=%s", string(msg))
	}

	go jc.ProcessQueue(string(rabbit.HatsQueue), f)

}
