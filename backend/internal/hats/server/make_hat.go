package server

import (
	"backend/internal/hats/repo"
	"backend/rpc/hatspb"
	"context"
	"math/rand"

	"github.com/sirupsen/logrus"
)

// MakeHat makes a hat
func (s *Server) MakeHat(ctx context.Context, size *hatspb.Size) (*hatspb.Hat, error) {

	logrus.Debugf("MakeHat() size=%v", size)

	hr := repo.GetRepo(ctx)

	color := []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)]

	name := []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)]

	mod := &repo.HatMod{
		Color:  color,
		Name:   name,
		Inches: size.GetInches(),
	}

	err := hr.Save(mod)

	if err != nil {
		return nil, err
	}

	return &hatspb.Hat{
		Color:  mod.Color,
		Name:   mod.Name,
		Inches: mod.Inches,
	}, nil
}
