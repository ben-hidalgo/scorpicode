package hatserver

import (
	"backend/internal/hats/repo"
	"backend/pkg/token"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// ListHats returns a list of hats
func (hs *Server) ListHats(ctx context.Context, req *hatspb.ListHatsRequest) (*hatspb.ListHatsResponse, error) {

	logrus.Debugf("ListHats() req=%v", req)

	// headers := httpwrap.GetHeaders(ctx)

	// logrus.Debugf("headers=%#v", headers)

	b := token.GetBearer(ctx)

	logrus.Debugf("ListHats() bearer=%#v", b)

	hr := repo.GetRepo(ctx)

	mods, err := hr.FindAll(repo.Limit(10), repo.Offset(0))
	if err != nil {
		return nil, err
	}

	hats := make([]*hatspb.Hat, len(mods))

	for i, m := range mods {
		hats[i] = &hatspb.Hat{
			Id:      m.ID,
			Version: int32(m.Version),
			Color:   m.Color,
			Style:   ToStyle(m.Style),
			Inches:  m.Inches,
		}
	}

	return &hatspb.ListHatsResponse{
		Hats: hats,
	}, nil
}
