package hatserver

import (
	"backend/internal/hats/repo"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
)

// FetchHat gets a hat by ID
func (hs *Server) FetchHat(ctx context.Context, req *hatspb.FetchHatRequest) (*hatspb.FetchHatResponse, error) {

	logrus.Debugf("FetchHat() req=%v", req)

	hr := repo.GetRepo(ctx)
	defer hr.Discard()

	mod, err := hr.Find(req.GetId())
	if err == repo.ErrNotFound {
		// TODO: should this be wrapped in util?
		// TODO: should find() return nil, nil rather than an err?
		return nil, twirp.NotFoundError(req.GetId())
	}
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}

	return &hatspb.FetchHatResponse{

		Hat: &hatspb.Hat{
			Id:      mod.ID,
			Color:   mod.Color,
			Style:   ToStyle(mod.Style),
			Size:    mod.Size,
			Version: int32(mod.Version),
		},
	}, nil
}
