package inmem

import (
	"backend/internal/hats/repo"
	"context"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

// Repo implements repo.HatRepo
type Repo struct {
	storage map[string]*repo.HatMod
}

// enforces the interface is implemented
var _ repo.HatRepo = (*Repo)(nil)

// NewRepo .
func NewRepo() *Repo {
	return &Repo{
		storage: make(map[string]*repo.HatMod),
	}
}

// FindAll .
func (r *Repo) FindAll(limit repo.Limit, offset repo.Offset) (hats []*repo.HatMod, err error) {
	// TODO: respect limit and offset
	for _, s := range r.storage {
		hats = append(hats, s)
	}
	return hats, nil
}

// Save .
func (r *Repo) Save(hm repo.HatMod) (string, error) {

	var id string
	if hm.ID == "" {
		id = xid.New().String()
	} else {
		id = hm.ID
	}

	mod := &repo.HatMod{
		ID:      id,
		Color:   hm.Color,
		Name:    hm.Name,
		Inches:  hm.Inches,
		Version: hm.Version + 1,
	}

	r.storage[id] = mod
	return id, nil
}

// Exists .
func (r *Repo) Exists(id string) (bool, error) {
	_, ok := r.storage[id]
	return ok, nil
}

// Delete .
func (r *Repo) Delete(id string, version int) error {
	v, ok := r.storage[id]

	if !ok {
		return repo.ErrNotFound
	}

	_ = v
	//TODO: re-add
	// if v.Version != version {
	// 	return repo.ErrVersionMismatch
	// }

	delete(r.storage, id)
	return nil
}

// Find .
func (r *Repo) Find(id string) (*repo.HatMod, error) {
	v, ok := r.storage[id]
	if ok {
		return v, nil
	}
	return nil, nil
}

////////// connection related

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
	r.storage = nil
	return nil
}
