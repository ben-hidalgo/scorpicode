package hatserver

import (
	"backend/internal/hats/repo"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"

	"github.com/twitchtv/twirp"

	"github.com/sirupsen/logrus"
)

var colors = map[string]interface{}{
	"RED":    struct{}{},
	"BLUE":   struct{}{},
	"GREEN":  struct{}{},
	"YELLOW": struct{}{},
	"PURPLE": struct{}{},
	"BLACK":  struct{}{},
	"GREY":   struct{}{},
	"ORANGE": struct{}{},
}

// MakeHat makes a hat
func (hs *Server) MakeHat(ctx context.Context, req *hatspb.MakeHatRequest) (*hatspb.MakeHatResponse, error) {

	logrus.Debugf("MakeHat() req=%v", req)

	if req.GetColor() == "" {
		return nil, util.InvalidArgumentError(HatColorRequired)
	}

	if req.GetSize() == "" {
		return nil, util.InvalidArgumentError(HatSizeRequired)
	}

	if _, ok := colors[req.GetColor()]; ok == false {
		return nil, util.InvalidArgumentError(HatColorDomain)
	}

	if req.GetStyle() == hatspb.Style_UNKNOWN_STYLE {
		return nil, util.InvalidArgumentError(HatStyleRequired)
	}

	// TODO: validate size slug

	hr := repo.GetRepo(ctx)
	if err := hr.Multi(); err != nil {
		return nil, util.InternalErrorWith(err)
	}
	defer hr.Discard()

	// a different instance is returned
	mod, err := hr.Save(repo.HatMod{
		Color: req.GetColor(),
		Style: req.GetStyle().String(),
		Size:  req.GetSize(),
	})
	if err != nil {
		return nil, err
	}

	if err := hr.Exec(); err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return &hatspb.MakeHatResponse{

		Hat: &hatspb.Hat{
			Id:      mod.ID,
			Color:   mod.Color,
			Style:   ToStyle(mod.Style),
			Size:    mod.Size,
			Version: int32(mod.Version),
		},
	}, nil
}
