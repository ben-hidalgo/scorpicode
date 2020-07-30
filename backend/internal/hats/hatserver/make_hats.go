package hatserver

import (
	"backend/internal/hats/orderdao"
	"backend/pkg/authnz"
	"backend/pkg/rabbit"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// MakeHats makes a hat
func (hs *Server) MakeHats(ctx context.Context, req *hatspb.MakeHatsRequest) (*hatspb.MakeHatsResponse, error) {

	logrus.Debugf("MakeHats() req=%#v", req)

	bearer := authnz.GetBearer(ctx)
	// only haberdasher role is allowed to make hats
	if !bearer.HasRole(authnz.HABERDASHER) && req.GetQuantity() > 100 {
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
		Color:     req.GetColor(),
		Style:     req.GetStyle(),
		Size:      req.GetSize(),
		Notes:     req.GetNotes(),
		Quantity:  req.GetQuantity(),
		CreatedBy: bearer.GetSubject(),
	}
	err := dao.Create(ctx, order)
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}

	rmq := rabbit.From(ctx)

	// publish the message to RabbitMQ
	rmq.SendJSON(rabbit.ServiceMsgActionX, rabbit.HatsOrderCreatedK, order)
	// actual hat creation is asynch in hatrabbit.ProcessHatsOrderCreated()

	res := &hatspb.MakeHatsResponse{
		Order: OrderDocToRep(order),
	}

	logrus.Debugf("MakeHats() res=%#v", res)

	return res, nil
}
