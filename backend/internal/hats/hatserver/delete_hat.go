package hatserver

import (
	"backend/internal/hats/repo"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// DeleteHat deletes a hat
func (hs *Server) DeleteHat(ctx context.Context, req *hatspb.DeleteHatRequest) (*hatspb.DeleteHatResponse, error) {

	logrus.Debugf("DeleteHat() req=%v", req)

	if req.GetId() == "" {
		return nil, util.InvalidArgumentError(HatIDRequired)
	}

	if req.GetVersion() == 0 {
		return nil, util.InvalidArgumentError(HatVersionRequired)
	}

	hr := repo.GetRepo(ctx)
	// if err := hr.Multi(); err != nil {
	// 	return nil, util.InternalErrorWith(err)
	// }
	// defer hr.Discard()

	err := hr.Delete(req.GetId(), int(req.GetVersion()))
	if err != nil {
		return nil, err
	}

	// if err := hr.Exec(); err != nil {
	// 	return nil, twirp.InternalErrorWith(err)
	// }

	return &hatspb.DeleteHatResponse{}, nil
}
