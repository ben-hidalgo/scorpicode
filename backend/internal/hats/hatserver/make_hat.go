package hatserver

import (
	"backend/internal/hats/repo"
	"backend/rpc/hatspb"
	"context"
	"math/rand"

	"github.com/twitchtv/twirp"

	"github.com/sirupsen/logrus"
)

// MakeHat makes a hat
func (hs *Server) MakeHat(ctx context.Context, req *hatspb.MakeHatRequest) (*hatspb.MakeHatResponse, error) {

	logrus.Debugf("MakeHat() req=%v", req)

	if req.GetInches() <= 0 {
		return nil, twirp.InvalidArgumentError(Inches, MustBeGTZero)
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

	mod := &repo.HatMod{
		Color:  color,
		Name:   name,
		Inches: req.GetInches(),
	}

	err := hr.Save(mod)
	if err != nil {
		return nil, err
	}

	err = hr.Commit()
	if err != nil {
		return nil, err
	}

	return &hatspb.MakeHatResponse{

		Hat: &hatspb.Hat{
			Color:  mod.Color,
			Name:   mod.Name,
			Inches: mod.Inches,
		},
	}, nil
}
