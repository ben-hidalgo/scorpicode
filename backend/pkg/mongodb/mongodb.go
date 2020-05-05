package mongodb

import (
	"backend/pkg/envconfig"
	"context"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseName .
var DatabaseName = "scdata"

// MongoURI .
var MongoURI = "mongodb://scuser:scpass@localhost:27017/scdata"

func init() {
	envconfig.SetString("MONGO_DB", &DatabaseName)
	envconfig.SetString("MONGO_URI", &MongoURI)
}

// Client returns the mgm client
func Client() (*mongo.Client, error) {

	err := mgm.SetDefaultConfig(nil, DatabaseName, options.Client().ApplyURI(MongoURI))
	if err != nil {
		return nil, err
	}

	// Get default connection client
	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
