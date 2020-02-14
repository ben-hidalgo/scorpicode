package hatserver

import (
	"backend/internal/hats/config"
	"backend/internal/hats/repo"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"
	"math/rand"

	"github.com/sirupsen/logrus"
)

// MakeHat makes a hat
func (hs *Server) MakeHat(ctx context.Context, req *hatspb.MakeHatRequest) (*hatspb.MakeHatResponse, error) {

	logrus.Debugf("MakeHat() req=%v", req)

	if req.GetInches() < config.MinSizeInches {
		return nil, util.InvalidArgumentError(Inches, HatTooSmall)
	}

	if req.GetInches() > config.MaxSizeInches {
		return nil, util.InvalidArgumentError(Inches, HatTooBig)
	}

	hr := repo.GetRepo(ctx)

	var color string
	if req.GetColor() == "" {
		color = []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)]
	} else {
		color = req.GetColor()
	}

	var name string
	if req.GetName() == "" {
		name = []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)]
	} else {
		name = req.GetName()
	}

	// a different instance is returned
	mod, err := hr.Save(repo.HatMod{
		Color:  color,
		Name:   name,
		Inches: req.GetInches(),
	})
	if err != nil {
		return nil, err
	}

	err = hr.Commit()
	if err != nil {
		return nil, err
	}

	return &hatspb.MakeHatResponse{

		Hat: &hatspb.Hat{
			Id:     mod.ID,
			Color:  mod.Color,
			Name:   mod.Name,
			Inches: mod.Inches,
		},
	}, nil
}
