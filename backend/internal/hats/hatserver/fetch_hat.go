package hatserver

import (
	"backend/internal/hats/hatdao"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// FetchHat gets a hat by ID
func (hs *Server) FetchHat(ctx context.Context, req *hatspb.FetchHatRequest) (*hatspb.FetchHatResponse, error) {

	logrus.Debugf("FetchHat() req=%v", req)

	hd := hatdao.From(ctx)

	hat, err := hd.Find(ctx, req.GetId())
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}
	if hat == nil {
		return nil, util.NotFoundError(req.GetId())
	}

	return &hatspb.FetchHatResponse{
		Hat: HatDocToRep(hat),
	}, nil
}
