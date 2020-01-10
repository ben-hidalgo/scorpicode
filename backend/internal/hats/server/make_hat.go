package server

import (
	"backend/rpc/hatspb"
	"context"
	"math/rand"
)

// MakeHat implements MakeHat procedure
func (s *Server) MakeHat(ctx context.Context, size *hatspb.Size) (hat *hatspb.Hat, err error) {

	return &hatspb.Hat{
		Inches: size.Inches,
		Color:  []string{"white", "black", "brown", "red", "blue"}[rand.Intn(4)],
		Name:   []string{"bowler", "baseball cap", "top hat", "derby"}[rand.Intn(3)],
	}, nil
}
