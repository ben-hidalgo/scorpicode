package server

import (
	"backend/internal/hats/repo"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// ListHats returns a list of hats
func (s *Server) ListHats(ctx context.Context, req *hatspb.ListHatsRequest) (*hatspb.ListHatsResponse, error) {

	logrus.Debugf("ListHats() req=%v", req)

	hr := repo.GetRepo(ctx)

	mods, err := hr.FindAll(repo.Limit(10), repo.Offset(0))
	if err != nil {
		return nil, err
	}

	hats := make([]*hatspb.Hat, len(mods))

	for i, m := range mods {
		hats[i] = &hatspb.Hat{
			Color:  m.Color,
			Name:   m.Name,
			Inches: m.Inches,
		}
	}

	return &hatspb.ListHatsResponse{
		Hats: hats,
	}, nil
}
