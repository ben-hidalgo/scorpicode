package orderdao

import (
	"context"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// Order .
type Order struct {
	// DefaultModel includes: _id,created_at and updated_at
	mgm.DefaultModel `bson:",inline"`
	Size             string `json:"size"          bson:"size"`
	Color            string `json:"color"         bson:"color"`
	Style            string `json:"style"         bson:"style"`
	Version          int32  `json:"version"       bson:"version"`
	Batch            string `json:"batch"       bson:"batch"`
	Quantity         int32  `json:"quantity"       bson:"quantity"`
	Notes            string `json:"notes"       bson:"notes"`
}

// OrderDao Order Data Access Object
type OrderDao interface {
	Create(context.Context, *Order) error
	Find(ctx context.Context, id string) (*Order, error) // not found returns nil, nil
}

// enforces the interface is implemented
var _ OrderDao = (*impl)(nil)

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
func New(mc *mongo.Client) OrderDao {
	return &impl{
		Client: mc,
	}
}

// From returns the dao and panics if not found
func From(ctx context.Context) OrderDao {

	switch v := ctx.Value(Key).(type) {
	case OrderDao:
		return v
	default:
		panic("orderdao.From() no value found")
	}
}
