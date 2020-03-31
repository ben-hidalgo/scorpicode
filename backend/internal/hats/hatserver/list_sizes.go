package hatserver

import (
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// ListSizes returns a list of hats
func (hs *Server) ListSizes(ctx context.Context, req *hatspb.ListSizesRequest) (*hatspb.ListSizesResponse, error) {

	logrus.Debugf("ListSizes() req=%v", req)

	sizes := []*hatspb.Size{
		&hatspb.Size{
			Slug: "06000",
			Name: "6",
		},
		&hatspb.Size{
			Slug: "06125",
			Name: "6 1/8",
		},
		&hatspb.Size{
			Slug: "06250",
			Name: "6 1/2",
		},
		&hatspb.Size{
			Slug: "06375",
			Name: "6 3/8",
		},
		&hatspb.Size{
			Slug: "06500",
			Name: "6 1/2",
		},
		&hatspb.Size{
			Slug: "06675",
			Name: "6 5/8",
		},
		&hatspb.Size{
			Slug: "06750",
			Name: "6 3/4",
		},
		&hatspb.Size{
			Slug: "06875",
			Name: "6 7/8",
		},
		&hatspb.Size{
			Slug: "0700",
			Name: "7",
		},
	}

	return &hatspb.ListSizesResponse{
		Sizes: sizes,
	}, nil
}
