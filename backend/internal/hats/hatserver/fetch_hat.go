package hatserver

import (
	"backend/internal/hats/hatsrepo"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// FetchHat gets a hat by ID
func (hs *Server) FetchHat(ctx context.Context, req *hatspb.FetchHatRequest) (*hatspb.FetchHatResponse, error) {

	logrus.Debugf("FetchHat() req=%v", req)

	hr := hatsrepo.FromContext(ctx)

	mod, err := hr.FindOneMakeHatsCmd(req.GetId())
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}
	if mod == nil {
		return nil, util.NotFoundError(req.GetId())
	}

	return &hatspb.FetchHatResponse{
		Hat: MakeHatsCmdToHat(mod),
	}, nil
}
