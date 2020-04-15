package mongorepo

import (
	"backend/internal/hats/hatsrepo"
	"backend/pkg/envconfig"
	"context"
	"fmt"

	"github.com/Kamva/mgm/v2"
	"github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
	"go.mongodb.org/mongo-driver/bson"
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

//MongoRepo implements HatsRepo
type MongoRepo struct {
}

// enforces the interface is implemented
var _ hatsrepo.HatsRepo = (*MongoRepo)(nil)

func init() {

	if err := mgm.SetDefaultConfig(nil, DatabaseName, options.Client().ApplyURI(MongoURI)); err != nil {
		panic(fmt.Sprintf("%#v", err))
	}
}

// NewRepo .
func NewRepo() *MongoRepo {
	return &MongoRepo{}
}

// ServerHooks is a Twirp middleware
func ServerHooks() *twirp.ServerHooks {

	// Ping the DB once to confirm connectivity
	if err := mgm.Coll(&hatsrepo.Hat{}).Database().Client().Ping(context.Background(), nil); err != nil {
		panic(fmt.Sprintf("%#v", err))
	}
	logrus.Debug("mongorepo.NewRepo() ping successful")

	return &twirp.ServerHooks{
		RequestReceived: func(ctx context.Context) (context.Context, error) {
			return context.WithValue(ctx, hatsrepo.RepoKey, NewRepo()), nil
		},
	}
}

// CreateHat .
func (r *MongoRepo) CreateHat(h *hatsrepo.Hat) error {
	h.Version = 1
	return mgm.Coll(h).Create(h)
}

// CreateMakeHatsCmd .
func (r *MongoRepo) CreateMakeHatsCmd(m *hatsrepo.MakeHatsCmd) error {
	m.Version = 1
	return mgm.Coll(m).Create(m)
}

// DeleteMakeHatsCmd calls the injected function
func (r *MongoRepo) DeleteMakeHatsCmd(id string, version int32) error {

	mhc := &hatsrepo.MakeHatsCmd{}

	coll := mgm.Coll(mhc)

	// TODO: validate not found behavior
	err := coll.FindByID(id, mhc)
	if err != nil {
		return err
	}

	if mhc.Version != version {
		return hatsrepo.ErrVersionMismatch
	}

	return coll.Delete(mhc)
}

// FindAllMakeHatsCmd .
func (r *MongoRepo) FindAllMakeHatsCmd() ([]*hatsrepo.MakeHatsCmd, error) {

	// result := []Book{}
	// err := mgm.Coll(&Book{}).SimpleFind(&result, bson.M{"age": bson.M{operator.Gt: 24}})

	results := []*hatsrepo.MakeHatsCmd{}

	err := mgm.Coll(&hatsrepo.MakeHatsCmd{}).SimpleFind(&results, bson.M{})

	if err != nil {
		return nil, err
	}

	return results, nil
}
