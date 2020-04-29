package hatserver

import (
	"backend/internal/hats/hatdao"
	"backend/pkg/authnz"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// ListHats returns a list of hats
func (hs *Server) ListHats(ctx context.Context, req *hatspb.ListHatsRequest) (*hatspb.ListHatsResponse, error) {

	logrus.Debugf("ListHats() req=%#v", req)

	b := authnz.GetBearer(ctx)

	logrus.Debugf("ListHats() bearer=%#v", b)

	hd := hatdao.From(ctx)

	docs, err := hd.Query(ctx)
	if err != nil {
		return nil, err
	}

	reps := make([]*hatspb.Hat, len(docs))

	for i, m := range docs {
		reps[i] = HatDocToRep(m)
	}

	// TODO: an empty array is not returned in JSON...
	return &hatspb.ListHatsResponse{
		Hats: reps,
	}, nil
}
