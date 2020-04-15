package hatserver

import (
	"backend/internal/hats/hatsrepo"
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

	hr := hatsrepo.FromContext(ctx)

	err := hr.DeleteMakeHatsCmd(req.GetId(), req.GetVersion())

	// TODO: ensure not found returns 404

	if err != nil {
		return nil, err
	}

	return &hatspb.DeleteHatResponse{}, nil
}
