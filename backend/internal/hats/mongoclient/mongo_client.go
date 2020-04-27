package mongoclient

import (
	"backend/internal/hats/hatdao"
	"backend/pkg/envconfig"
	"context"
	"fmt"

	"github.com/Kamva/mgm/v2"
	"github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseName .
var DatabaseName = "hats"

// MongoURI . TODO: use env var
var MongoURI = "mongodb://hats:hats@localhost:27017/hats"

func init() {
	envconfig.SetString("MONGO_DB", &DatabaseName)
	envconfig.SetString("MONGO_URI", &MongoURI)
}

func init() {

	if err := mgm.SetDefaultConfig(nil, DatabaseName, options.Client().ApplyURI(MongoURI)); err != nil {
		panic(fmt.Sprintf("%#v", err))
	}
}

// ServerHooks is a Twirp middleware
func ServerHooks() *twirp.ServerHooks {

	// Ping the DB once to confirm connectivity
	if err := mgm.Coll(&hatdao.Hat{}).Database().Client().Ping(context.Background(), nil); err != nil {
		panic(fmt.Sprintf("%#v", err))
	}
	logrus.Debug("mongorepo.NewRepo() ping successful")

	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {

			// each Dao will be added to the context separately
			ctx = context.WithValue(ctx, hatdao.Key, hatdao.New())

			return ctx, nil
		},
	}
}
