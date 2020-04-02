package mongorepo

import (
	"backend/internal/hats/hatsrepo"
	"backend/pkg/envconfig"

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

// Init .
func Init() error {
	// Setup mgm default config
	err := mgm.SetDefaultConfig(nil, DatabaseName, options.Client().ApplyURI(MongoURI))
	return err
}

//MongoRepo implements HatsRepo
type MongoRepo struct {
}

// enforces the interface is implemented
var _ hatsrepo.HatsRepo = (*MongoRepo)(nil)

// Save .
func (r *MongoRepo) Save(b *hatsrepo.Book) (*mgm.IDField, error) {
	return nil, nil
}
