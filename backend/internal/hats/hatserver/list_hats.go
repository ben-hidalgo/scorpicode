package hatserver

import (
	"backend/internal/hats/hatsrepo"
	"backend/pkg/authnz"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// ListHats returns a list of hats
func (hs *Server) ListHats(ctx context.Context, req *hatspb.ListHatsRequest) (*hatspb.ListHatsResponse, error) {

	logrus.Debugf("ListHats() req=%#v", req)

	// headers := httpwrap.GetHeaders(ctx)

	// logrus.Debugf("headers=%#v", headers)

	b := authnz.GetBearer(ctx)

	logrus.Debugf("ListHats() bearer=%#v", b)

	logrus.Debugf("ListHats() b.GetEmail()=%s", b.GetEmail())
	logrus.Debugf("ListHats() b.GetRoles()=%s", b.GetRoles())

	hr := hatsrepo.FromContext(ctx)

	mods, err := hr.FindAllMakeHatsCmd()
	if err != nil {
		return nil, err
	}

	hats := make([]*hatspb.Hat, len(mods))

	for i, m := range mods {
		hats[i] = MakeHatsCmdToHat(m)
	}

	return &hatspb.ListHatsResponse{
		Hats: hats,
	}, nil
}
