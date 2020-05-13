package hatserver

import (
	"backend/internal/hats/orderdao"
	"backend/pkg/util"
	"backend/rpc/hatspb"
	"context"

	"github.com/sirupsen/logrus"
)

// FetchOrder gets an order by ID
func (hs *Server) FetchOrder(ctx context.Context, req *hatspb.FetchOrderRequest) (*hatspb.FetchOrderResponse, error) {

	logrus.Debugf("FetchOrder() req=%v", req)

	od := orderdao.From(ctx)

	order, err := od.Find(ctx, req.GetId())
	if err != nil {
		return nil, util.InternalErrorWith(err)
	}
	if order == nil {
		return nil, util.NotFoundError(req.GetId())
	}

	return &hatspb.FetchOrderResponse{
		Order: OrderDocToRep(order),
	}, nil
}
