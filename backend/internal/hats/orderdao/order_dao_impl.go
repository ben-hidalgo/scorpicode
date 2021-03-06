package orderdao

import (
	"context"
	"encoding/hex"

	"github.com/Kamva/mgm/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create .
func (i *impl) Create(ctx context.Context, o *Order) error {
	o.Version = 1
	return mgm.Coll(o).CreateWithCtx(ctx, o)
}

// Find not found returns nil, nil
func (i *impl) Find(ctx context.Context, id string) (*Order, error) {

	o := &Order{}

	coll := mgm.Coll(o)

	err := coll.FindByIDWithCtx(ctx, id, o)
	if err == hex.ErrLength || err == mongo.ErrNoDocuments {
		logrus.Warnf("orderdao.Find() warn err=%s", err)
		// malformed id means not found
		return nil, nil
	}
	if err != nil {
		logrus.Errorf("orderdao.Find() err=%s", err)
		return nil, err
	}
	return o, nil
}
