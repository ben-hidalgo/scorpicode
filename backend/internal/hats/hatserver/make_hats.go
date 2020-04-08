package hatserver

import (
	"backend/internal/hats/hatsrepo"
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

	// TODO: validate size slug, quantity and notes

	// TODO: transaction func
	hr := hatsrepo.FromContext(ctx)

	cmd := &hatsrepo.MakeHatsCmd{
		Color:    req.GetColor(),
		Style:    req.GetStyle().String(),
		Size:     req.GetSize(),
		Quantity: req.GetQuantity(),
		Notes:    req.GetNotes(),
	}

	// the passed-in cmd will be mutated
	err := hr.CreateMakeHatsCmd(cmd)
	if err != nil {
		return nil, err
	}

	logrus.Debugf("MakeHats() cmd=%#v", cmd)

	// TODO: save a hat for each quantity with foreign key to the cmd

	res := &hatspb.MakeHatsResponse{

		Hat: &hatspb.Hat{
			Id:       cmd.ID.Hex(),
			Color:    cmd.Color,
			Style:    ToStyle(cmd.Style),
			Size:     cmd.Size,
			Quantity: int32(cmd.Quantity),
			Version:  int32(cmd.Version),
			Notes:    cmd.Notes,
		},
	}

	logrus.Debugf("MakeHats() res=%#v", res)

	return res, nil
}
