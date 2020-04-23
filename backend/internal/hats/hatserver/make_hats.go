package hatserver

import (
	"backend/internal/hats/hatsrepo"
	"backend/pkg/authnz"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

var colors = map[string]interface{}{
	"RED":    struct{}{},
	"BLUE":   struct{}{},
	"GREEN":  struct{}{},
	"YELLOW": struct{}{},
	"PURPLE": struct{}{},
	"BLACK":  struct{}{},
	"GREY":   struct{}{},
	"ORANGE": struct{}{},
}

// MakeHats makes a hat
func (hs *Server) MakeHats(ctx context.Context, req *hatspb.MakeHatsRequest) (*hatspb.MakeHatsResponse, error) {

	logrus.Debugf("MakeHats() req=%#v", req)

	bearer := authnz.GetBearer(ctx)
	// only haberdasher role is allowed to make hats
	if !bearer.HasRole(authnz.HABERDASHER) {
		return nil, util.PermissionDeniedError(MakeHatsForbidden)
	}

	if req.GetColor() == "" {
		return nil, util.InvalidArgumentError(HatColorRequired)
	}

	if req.GetSize() == "" {
		return nil, util.InvalidArgumentError(HatSizeRequired)
	}

	if _, ok := colors[req.GetColor()]; ok == false {
		return nil, util.InvalidArgumentError(HatColorDomain)
	}

	if req.GetStyle() == hatspb.Style_UNKNOWN_STYLE {
		return nil, util.InvalidArgumentError(HatStyleRequired)
	}

	if req.GetQuantity() <= 0 {
		return nil, util.InvalidArgumentError(HatQuantityInvalid)
	}

	// TODO: validate size slug, quantity and notes

	// TODO: transaction func
	hr := hatsrepo.FromContext(ctx)

	// TODO: save the created by user id
	cmd := &hatsrepo.MakeHatsCmd{
		Color:    req.GetColor(),
		Style:    req.GetStyle().String(),
		Size:     req.GetSize(),
		Quantity: req.GetQuantity(),
		Notes:    req.GetNotes(),
	}

	tf := func() error {
		// the passed-in cmd will be mutated
		err := hr.CreateMakeHatsCmd(cmd)
		if err != nil {
			return err
		}

		// TODO: save a hat for each quantity with foreign key to the cmd

		for i := int32(0); i < cmd.Quantity; i++ {
			h := &hatsrepo.Hat{
				// TODO: correct datatype for MakeHatsCmdID?
				MakeHatsCmdID: cmd.ID.Hex(),
				Color:         cmd.Color,
				Style:         cmd.Style,
				Size:          cmd.Size,
				// quantity and notes are MakeHatsCmd level only
			}
			err := hr.CreateHat(h)
			if err != nil {
				return err
			}
		}

		return nil
	}

	hr.VisitTxn(ctx, tf)

	// reusable for list hats
	res := &hatspb.MakeHatsResponse{
		// TODO: rename "Hat" in the response
		Hat: MakeHatsCmdToHat(cmd),
	}

	logrus.Debugf("MakeHats() res=%#v", res)

	return res, nil
}
