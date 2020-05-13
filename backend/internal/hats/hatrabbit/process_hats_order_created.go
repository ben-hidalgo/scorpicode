package hatrabbit

import (
	"backend/internal/hats/hatdao"
	"backend/internal/hats/orderdao"
	"backend/pkg/rabbit"
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// ProcessHatsOrderCreated .
func ProcessHatsOrderCreated(ctx context.Context, msg []byte) error {

	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() msg=%s", string(msg))

	hd := hatdao.From(ctx)
	od := orderdao.From(ctx)
	rmq := rabbit.From(ctx)

	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() hd=%#v", hd)
	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() od=%#v", od)
	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() rmq=%#v", rmq)

	var order orderdao.Order
	err := json.Unmarshal(msg, &order)
	if err != nil {
		return err
	}
	logrus.Infof("hatrabbit.ProcessHatsOrderCreated() order=%#v", order)

	// this is the transaction func
	tf := func() error {

		// save a hat per quantity
		for i := int32(0); i < order.Quantity; i++ {
			hat := &hatdao.Hat{
				Color:   order.Color,
				Style:   order.Style,
				Size:    order.Size,
				Ordinal: i,
				OrderID: order.ID,
			}
			err := hd.Create(ctx, hat)
			if err != nil {
				return err
			}

			rmq.SendJSON(rabbit.ServiceMsgActionX, rabbit.HatsHatCreatedK, hat)
		}
		return nil
	}

	err = hd.VisitTxn(ctx, tf)
	if err != nil {
		return err
	}

	return nil
}
