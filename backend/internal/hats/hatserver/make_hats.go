package hatserver

import (
	"backend/internal/hats/hatdao"
	"backend/pkg/authnz"
	"backend/pkg/rabbit"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"
	"time"

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

	hd := hatdao.From(ctx)

	// TODO: save the created by user id

	// TODO: used for publishing RMQ message
	var docs []*hatdao.Hat

	// use a transaction so all hats or none are committed
	tf := func() error {

		// save a hat per quantity
		for i := int32(0); i < req.GetQuantity(); i++ {
			hat := &hatdao.Hat{
				Color: req.GetColor(),
				Style: req.GetStyle().String(),
				Size:  req.GetSize(),
				// TODO: add notes
			}
			err := hd.Create(ctx, hat)
			if err != nil {
				return err
			}
			docs = append(docs, hat)
		}

		return nil
	}

	// TODO: add test to ensure the internal error is returned
	err := hd.VisitTxn(ctx, tf)
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}

	rmq := rabbit.From(ctx)

	logrus.Infof("rmq=%#v", rmq)

	// TODO: add "envelope" message type in proto
	// publish the message
	rmq.SendJSON(rabbit.ServiceMsgtypeTx, rabbit.HatsDotMakeHats, docs)

	res := &hatspb.MakeHatsResponse{
		// TODO: modify response structure?
		Hat: HatDocToRep(docs[0]),
	}

	logrus.Debugf("MakeHats() res=%#v", res)

	return res, nil
}

// HatDocToRep convert Hat document (Mongo) to Hat representation (gRPC)
func HatDocToRep(hat *hatdao.Hat) *hatspb.Hat {
	return &hatspb.Hat{
		Id:        hat.ID.Hex(),
		CreatedAt: hat.CreatedAt.Format(time.RFC3339),
		UpdatedAt: hat.UpdatedAt.Format(time.RFC3339),
		Version:   int32(hat.Version),
		Color:     hat.Color,
		Style:     ToStyle(hat.Style),
		Size:      hat.Size,
		// Quantity:  int32(hat.Quantity),
		// TODO: add notes to mod
		// Notes:     hat.Notes,
	}
}
