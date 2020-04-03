package mongorepo

import (
	"backend/internal/hats/hatsrepo"
	"backend/pkg/envconfig"
	"fmt"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseName .
var DatabaseName = "hats"

// MongoURI . TODO: use env var
var MongoURI = "mongodb://hats:hats@localhost:27017/hats"

func init() {
	envconfig.SetString("DATABASE_NAME", &DatabaseName)
	envconfig.SetString("MONGO_URI", &MongoURI)
}

//MongoRepo implements HatsRepo
type MongoRepo struct {
}

// enforces the interface is implemented
var _ hatsrepo.HatsRepo = (*MongoRepo)(nil)

var initialized = false

// NewRepo .
func NewRepo() *MongoRepo {
	if !initialized {
		if err := mgm.SetDefaultConfig(nil, DatabaseName, options.Client().ApplyURI(MongoURI)); err != nil {
			panic(fmt.Sprintf("%#v", err))
		}
		initialized = true
	}
	return &MongoRepo{}
}

// SaveHat .
func (r *MongoRepo) SaveHat(h *hatsrepo.Hat) error {

	err := mgm.Coll(h).Create(h)
	if err != nil {
		return err
	}

	return nil
}
