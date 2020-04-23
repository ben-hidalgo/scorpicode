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

	mhc, err := hr.FindOneMakeHatsCmd(ctx, req.GetId())
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}
	if mhc == nil {
		return nil, util.NotFoundError(req.GetId())
	}

	return &hatspb.FetchHatResponse{
		Hat: MakeHatsCmdToHat(mhc),
	}, nil
}
