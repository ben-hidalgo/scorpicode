package hatserver

import (
	"backend/internal/hats/hatdao"
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

	hd := hatdao.From(ctx)

	hat, err := hd.Find(ctx, req.GetId())
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}
	if hat == nil {
		return nil, util.NotFoundError(fmt.Sprintf("id: '%s', version: %d", req.GetId(), req.GetVersion()))
	}
	if hat.Version != req.GetVersion() {
		return nil, util.NotFoundError(fmt.Sprintf("id: '%s' version mismatch", req.GetId()))
	}

	// modify the signature to accept the found mhc
	err = hd.Delete(ctx, hat)
	if err != nil {
		return nil, err
	}

	return &hatspb.DeleteHatResponse{}, nil
}
