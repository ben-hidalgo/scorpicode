package hatserver

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
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

var styles = map[string]interface{}{
	"BOWLER":   struct{}{},
	"FEDORA":   struct{}{},
	"BASEBALL": struct{}{},
	"NEWSBOY":  struct{}{},
	"COWBOY":   struct{}{},
	"DERBY":    struct{}{},
	"TOP_HAT":  struct{}{},
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
		return nil, util.InvalidArgumentError(HatColorInvalid)
	}

	if req.GetStyle() == "" {
		return nil, util.InvalidArgumentError(HatStyleRequired)
	}

	if _, ok := styles[req.GetStyle()]; ok == false {
		return nil, util.InvalidArgumentError(HatStyleInvalid)
	}

	if req.GetQuantity() <= 0 {
		return nil, util.InvalidArgumentError(HatQuantityInvalid)
	}

	dao := orderdao.From(ctx)

	order := &orderdao.Order{
		Color:    req.GetColor(),
		Style:    req.GetStyle(),
		Size:     req.GetSize(),
		Notes:    req.GetNotes(),
		Quantity: req.GetQuantity(),
		// TODO: use subject
		CreatedBy: bearer.GetEmail(),
	}
	err := dao.Create(ctx, order)
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}

	rmq := rabbit.From(ctx)

	logrus.Infof("rmq=%#v", rmq)

	// TODO: add "envelope" message type in proto
	// publish the message
	rmq.SendJSON(rabbit.ServiceMsgtypeTx, rabbit.HatsDotMakeHats, order)

	res := &hatspb.MakeHatsResponse{
		Order: OrderDocToRep(order),
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
		Style:     hat.Style,
		Size:      hat.Size,
		// Quantity:  int32(hat.Quantity),
		// TODO: add notes to mod
		// Notes:     hat.Notes,
	}
}

// OrderDocToRep convert Order document (Mongo) to Order representation (gRPC)
func OrderDocToRep(order *orderdao.Order) *hatspb.Order {
	return &hatspb.Order{
		Id:        order.ID.Hex(),
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		UpdatedAt: order.UpdatedAt.Format(time.RFC3339),
		Version:   int32(order.Version),
		Color:     order.Color,
		Style:     order.Style,
		Size:      order.Size,
		Quantity:  int32(order.Quantity),
		Notes:     order.Notes,
	}
}
