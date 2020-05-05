package hatrabbit

import (
	"backend/pkg/rabbit"

	"github.com/sirupsen/logrus"
)

// Listen .
func Listen() {

	conn := rabbit.Connect()

	// Handler function
	f := func(msg []byte) {
		// TODO: use channels and switch by message type
		logrus.Infof("hatrabbit.Connect() msg=%s", string(msg))
	}

	go conn.ProcessQueue(string(rabbit.HatsQueue), f)

}
