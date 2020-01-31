package server

import (
	"backend/internal/hats/repo"
	"backend/rpc/hatspb"
	"context"
	"math/rand"

	"github.com/sirupsen/logrus"
)

// MakeHat makes a hat
func (s *Server) MakeHat(ctx context.Context, req *hatspb.MakeHatRequest) (*hatspb.MakeHatResponse, error) {

	logrus.Debugf("MakeHat() req=%v", req)

	hr := repo.GetRepo(ctx)

	color := []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)]

	name := []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)]

	mod := &repo.HatMod{
		Color:  color,
		Name:   name,
		Inches: req.GetInches(),
	}

	err := hr.Save(mod)

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
