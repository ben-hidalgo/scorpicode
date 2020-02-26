package hatserver

import (
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// ListSizes returns a list of hats
func (hs *Server) ListSizes(ctx context.Context, req *hatspb.ListSizesRequest) (*hatspb.ListSizesResponse, error) {

	logrus.Debugf("ListSizes() req=%v", req)

	inches := hatspb.Units(hatspb.Units_value["INCHES"])

	sizes := []*hatspb.Size{
		&hatspb.Size{
			Slug:      "06_00",
			Name:      "6 inches",
			Dimension: 6.0,
			Units:     inches,
		},
		&hatspb.Size{
			Slug:      "06_25",
			Name:      "6 1/4 inches",
			Dimension: 6.25,
			Units:     inches,
		},
		&hatspb.Size{
			Slug:      "06_50",
			Name:      "6 1/2 inches",
			Dimension: 6.5,
			Units:     inches,
		},
		&hatspb.Size{
			Slug:      "06_75",
			Name:      "6 3/4 inches",
			Dimension: 6.75,
			Units:     inches,
		},
		&hatspb.Size{
			Slug:      "07_00",
			Name:      "7 inches",
			Dimension: 7.0,
			Units:     inches,
		},
	}

	return &hatspb.ListSizesResponse{
		Sizes: sizes,
	}, nil
}
