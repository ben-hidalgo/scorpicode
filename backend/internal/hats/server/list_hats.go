package server

import (
	"backend/internal/hats/repo"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// ListHats returns a list of hats
func (s *Server) ListHats(ctx context.Context, req *hatspb.ListHatsRequest) (*hatspb.ListHatsResponse, error) {

	hr := repo.GetRepo(ctx)
	logrus.Debugf("ListHats() hr=%v", hr)

	fedora := &hatspb.Hat{
		Color:  "red",
		Name:   "fedora",
		Inches: 10,
	}

	bowler := &hatspb.Hat{
		Color:  "blue",
		Name:   "bowler",
		Inches: 12,
	}

	derby := &hatspb.Hat{
		Color:  "brown",
		Name:   "derby",
		Inches: 11,
	}

	hats := []*hatspb.Hat{
		fedora,
		bowler,
		derby,
	}

	return &hatspb.ListHatsResponse{
		Hats: hats,
	}, nil
}
