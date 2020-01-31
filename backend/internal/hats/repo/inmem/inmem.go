package inmem

import (
	"backend/internal/hats/repo"
	"context"

	"github.com/sirupsen/logrus"
)

//Repo implements repo.HatRepo
type Repo struct{}

// enforces the interface is implemented
var _ repo.HatRepo = (*Repo)(nil)

var storage = make(map[string]*repo.HatMod)

// NewRepo returns a pointer to a new instance of Repo
func NewRepo() *Repo {
	return &Repo{}
}

//BeginTx implements HatRepo.BeginTxn()
func (r *Repo) BeginTxn(ctx context.Context) error {
	logrus.Debug("inmem.BeginTxn()")
	return nil
}

//Rollback implements HatRepo.Rollback()
func (r *Repo) Rollback() error {
	logrus.Debug("inmem.Rollback()")
	return nil
}

//Commit implements HatRepo.Commit()
func (r *Repo) Commit() error {
	logrus.Debug("inmem.Commit()")
	return nil
}
