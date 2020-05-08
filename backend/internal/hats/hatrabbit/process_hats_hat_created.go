package hatrabbit

import (
	"context"

	"github.com/sirupsen/logrus"
)

// ProcessHatsHatCreated .
func ProcessHatsHatCreated(ctx context.Context, msg []byte) error {
	logrus.Infof("hatrabbit.ProcessHatsHatCreated() msg=%s", string(msg))
	return nil
}
