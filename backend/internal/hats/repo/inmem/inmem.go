package inmem

import (
	"backend/internal/hats/repo"
	"context"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

// Repo implements repo.HatRepo
type Repo struct{}

// enforces the interface is implemented
var _ repo.HatRepo = (*Repo)(nil)

var storage = make(map[string]*repo.HatMod)

// NewRepo returns a pointer to a new instance of Repo
func NewRepo() *Repo {
	return &Repo{}
}

// FindAll queries all records
func (r *Repo) FindAll(limit repo.Limit, offset repo.Offset) (hats []*repo.HatMod, err error) {
	// TODO: respect limit and offset
	for _, s := range storage {
		hats = append(hats, s)
	}
	return hats, nil
}

// Save performs an upsert, assigns an ID
func (r *Repo) Save(hm *repo.HatMod) error {
	if hm.ID == "" {
		hm.ID = xid.New().String()
	}
	storage[hm.ID] = hm
	return nil
}

// Exists returns true if the record exists
func (r *Repo) Exists(id string) (bool, error) {
	_, ok := storage[id]
	return ok, nil
}

// BeginTxn implements HatRepo.BeginTxn()
func (r *Repo) BeginTxn(ctx context.Context) error {
	logrus.Debug("inmem.BeginTxn()")
	return nil
}

// Rollback implements HatRepo.Rollback()
func (r *Repo) Rollback() error {
	logrus.Debug("inmem.Rollback()")
	return nil
}

// Commit implements HatRepo.Commit()
func (r *Repo) Commit() error {
	logrus.Debug("inmem.Commit()")
	return nil
}

// Close implements HatRepo.Close()
func (r *Repo) Close() error {
	logrus.Debug("inmem.Close()")
	return nil
}
