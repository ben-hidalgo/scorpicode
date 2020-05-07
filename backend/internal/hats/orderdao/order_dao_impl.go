package orderdao

import (
	"context"
	"encoding/hex"

	"github.com/Kamva/mgm/v2"
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
		// malformed id means not found
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return o, nil
}
