package hatserver

import (
	"backend/internal/hats/hatsrepo"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"
	"fmt"

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

	mhc, err := hr.FindOneMakeHatsCmd(ctx, req.GetId())
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}
	if mhc == nil {
		return nil, util.NotFoundError(fmt.Sprintf("id: '%s', version: %d", req.GetId(), req.GetVersion()))
	}
	if mhc.Version != req.GetVersion() {
		return nil, util.NotFoundError(fmt.Sprintf("id: '%s' version mismatch", req.GetId()))
	}

	// modify the signature to accept the found mhc
	err = hr.DeleteMakeHatsCmd(ctx, mhc)
	if err != nil {
		return nil, err
	}

	return &hatspb.DeleteHatResponse{}, nil
}
