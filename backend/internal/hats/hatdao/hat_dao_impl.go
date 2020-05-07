package hatdao

import (
	"context"
	"encoding/hex"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create .
func (i *impl) Create(ctx context.Context, e *Hat) error {
	e.Version = 1
	return mgm.Coll(e).CreateWithCtx(ctx, e)
}

// Delete .
func (i *impl) Delete(ctx context.Context, e *Hat) error {
	return mgm.Coll(e).DeleteWithCtx(ctx, e)
}

// Find not found returns nil, nil
func (i *impl) Find(ctx context.Context, id string) (*Hat, error) {

	h := &Hat{}

	coll := mgm.Coll(h)

	err := coll.FindByIDWithCtx(ctx, id, h)
	if err == hex.ErrLength || err == mongo.ErrNoDocuments {
		// malformed id means not found
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return h, nil
}

// Query .
func (i *impl) Query(ctx context.Context) ([]*Hat, error) {

	// result := []Book{}
	// err := mgm.Coll(&Book{}).SimpleFind(&result, bson.M{"age": bson.M{operator.Gt: 24}})

	results := []*Hat{}

	err := mgm.Coll(&Hat{}).SimpleFindWithCtx(ctx, &results, bson.M{})

	if err != nil {
		return nil, err
	}

	return results, nil
}

// VisitTxn .
func (i *impl) VisitTxn(ctx context.Context, tf func() error) error {

	return mgm.TransactionWithCtx(ctx, func(session mongo.Session, sc mongo.SessionContext) error {

		err := tf()
		if err != nil {
			return err
		}

		return session.CommitTransaction(sc)
	})
}
