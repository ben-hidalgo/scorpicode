package server

import (
	"backend/rpc/hatspb"
	"context"
)

// ListHats returns a list of hats
func (s *Server) ListHats(context.Context, *hatspb.ListHatsRequest) (*hatspb.ListHatsResponse, error) {

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
