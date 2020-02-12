package repo

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
)

// Offset is the page number
type Offset int

// Limit is the page size
type Limit int

// represents a Hat stored in the repo
type HatMod struct {
	ID     string
	Inches int32  `redis:"inches"`
	Color  string `redis:"color"`
	Name   string `redis:"name"`
}

//TODO: implement create and list in memory and update list_hats and make_hat
// the repository type
type HatRepo interface {
	BeginTxn(ctx context.Context) error
	Rollback() error
	Commit() error

	//FindAll queries all records
	FindAll(limit Limit, offset Offset) ([]*HatMod, error)
	//Save performs an upsert
	Save(hm *HatMod) error
}

// used to store the Repo in Context
type key int

// RepoKey is the key for the repo in context
var RepoKey key

//GetRepo returns the repo and panics if not found
func GetRepo(ctx context.Context) HatRepo {

	switch v := ctx.Value(RepoKey).(type) {
	case HatRepo:
		return v
	default:
		panic("GetRepo() no value found")
	}
}

// Hook middleware injects the DB impl
func Hook(hr HatRepo) *twirp.ServerHooks {

	hook := &twirp.ServerHooks{}

	hook.RequestReceived = func(ctx context.Context) (context.Context, error) {

		err := hr.BeginTxn(ctx)
		if err != nil {
			logrus.Errorf("repo.Hook() BeginTx failed err=%s", err)
			return ctx, err
		}

		return context.WithValue(ctx, RepoKey, hr), nil
	}

	hook.ResponseSent = func(ctx context.Context) {

		// call rollback to close the connection in case the procedure
		// returned before getting the repo (or forgot to defer Rollback)
		err := hr.Rollback()
		if err != nil {
			logrus.Errorf("repo.Hook() BeginTx failed err=%s", err)
		}
	}

	return hook
}
