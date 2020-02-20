package hatserver

import (
	"backend/internal/hats/config"
	"backend/internal/hats/repo"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"

	"github.com/twitchtv/twirp"

	"github.com/sirupsen/logrus"
)

// MakeHat makes a hat
func (hs *Server) MakeHat(ctx context.Context, req *hatspb.MakeHatRequest) (*hatspb.MakeHatResponse, error) {

	logrus.Debugf("MakeHat() req=%v", req)

	if req.GetInches() < config.MinSizeInches {
		return nil, util.InvalidArgumentError(HatInchesTooSmall)
	}

	if req.GetInches() > config.MaxSizeInches {
		return nil, util.InvalidArgumentError(HatInchesTooBig)
	}

	if req.GetColor() == "" {
		return nil, util.InvalidArgumentError(HatColorRequired)
	}

	if req.GetName() == "" {
		return nil, util.InvalidArgumentError(HatNameRequired)
	}

	hr := repo.GetRepo(ctx)
	if err := hr.Multi(); err != nil {
		return nil, twirp.InternalErrorWith(err)
	}
	defer hr.Discard()

	// a different instance is returned
	mod, err := hr.Save(repo.HatMod{
		Color:  req.GetColor(),
		Name:   req.GetName(),
		Inches: req.GetInches(),
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
			Name:    mod.Name,
			Inches:  mod.Inches,
			Version: int32(mod.Version),
		},
	}, nil
}
