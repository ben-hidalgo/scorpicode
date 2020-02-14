package repo

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
)

// Offset is the page number
type Offset int

// Limit is the page size
type Limit int

var (
	// ErrNotFound .
	ErrNotFound = errors.New("notfound")

	// ErrFieldMissing .
	ErrFieldMissing = errors.New("fieldmissing")

	// ErrVersionMismatch .
	ErrVersionMismatch = errors.New("versionmismatch")
)

// HatMod represents a Hat stored in the repo
type HatMod struct {
	ID      string
	Inches  int32  `redis:"inches"`
	Color   string `redis:"color"`
	Name    string `redis:"name"`
	Version int    `redis:"version"`
}

// HatRepo interface for data storage
type HatRepo interface {
	BeginTxn(ctx context.Context) error
	Rollback() error
	Commit() error
	Close() error

	// Find returns nil, nil if not found
	Find(id string) (*HatMod, error)

	// FindAll queries all records
	FindAll(limit Limit, offset Offset) ([]*HatMod, error)

	// Save performs an upsert, returns the ID
	// Input parameter is not mutated
	// Assigns an ID if not provided
	// Increments Version
	// Returns NotFound if missing by ID
	// Returns VersionMismatch if version isn't equal
	Save(hm HatMod) (*HatMod, error)

	// Exists returns true if the record exists
	Exists(id string) (bool, error)

	// Delete deletes the record if version matches;
	// throws NotFound, VersionMismatch
	Delete(id string, version int) error
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
