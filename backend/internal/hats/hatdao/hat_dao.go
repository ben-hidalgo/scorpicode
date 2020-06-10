package hatdao

import (
	"context"
	"encoding/hex"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Hat .
type Hat struct {
	// DefaultModel includes: _id,created_at and updated_at
	mgm.DefaultModel `bson:",inline"`
	Size             string             `json:"size"          bson:"size"`
	Color            string             `json:"color"         bson:"color"`
	Style            string             `json:"style"         bson:"style"`
	Version          int32              `json:"version"       bson:"version"`
	Ordinal          int32              `json:"ordinal"       bson:"ordinal"`
	OrderID          primitive.ObjectID `json:"order_id"      bson:"order_id"`
	CreatedBy        string             `json:"created_by"       bson:"created_by"`
}

// HatDao Hat Data Access Object
type HatDao interface {
	Create(context.Context, *Hat) error
	Delete(context.Context, *Hat) error
	Find(ctx context.Context, id string) (*Hat, error) // not found returns nil, nil
	Query(context.Context) ([]*Hat, error)
	VisitTxn(context.Context, func() error) error
}

// enforces the interface is implemented
var _ HatDao = (*impl)(nil)

// used to store the Repo in Context
type key int

// Key is the key for the repo in context; public for mock injection
const Key key = 0

//impl .
type impl struct {
	// mgm doesn't require an external connection to the client but this structure allows for libs or custom methods that would
	// and also allows the main method to defer close() and handle connection errors explicitly
	Client *mongo.Client
}

// New .
func New(mc *mongo.Client) HatDao {
	return &impl{
		Client: mc,
	}
}

// From returns the dao and panics if not found
func From(ctx context.Context) HatDao {

	switch v := ctx.Value(Key).(type) {
	case HatDao:
		return v
	default:
		panic("hatdao.From() no value found")
	}
}

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

// Query . TODO: add pagination parameters
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
