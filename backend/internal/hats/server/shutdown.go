package server

import (
	"backend/internal/hats/config"
	"backend/rpc/hatspb"
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

// MakeHat implements MakeHat procedure
func (s *Server) Shutdown(context.Context, *hatspb.ShutdownRequest) (*hatspb.ShutdownResponse, error) {
	logrus.Infof("%s shutting down", config.AppName)
	os.Exit(0)

	return &hatspb.ShutdownResponse{}, nil
}
