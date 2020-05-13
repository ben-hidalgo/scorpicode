package hatdao

import (
	"context"

	"github.com/Kamva/mgm/v2"
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
	// TODO: add batch UUID
	// TODO: add notes
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
