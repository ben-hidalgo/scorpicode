package hatserver

import (
	"backend/internal/hats/hatsrepo"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"
	"encoding/hex"

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

	// TODO: fetch mhc by id and version, return 404 if necessary

	// modify the signature to accept the found mhc
	err := hr.DeleteMakeHatsCmd(req.GetId(), req.GetVersion())
	if err == hex.ErrLength {
		return nil, util.NotFoundError(err.Error())
	}

	if err != nil {
		return nil, err
	}

	return &hatspb.DeleteHatResponse{}, nil
}
