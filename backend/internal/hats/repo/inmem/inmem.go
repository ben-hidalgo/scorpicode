package inmem

import (
	"backend/internal/hats/repo"

	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
)

// Repo implements repo.HatRepo
type Repo struct {
	Storage map[string]*repo.HatMod
}

// enforces the interface is implemented
var _ repo.HatRepo = (*Repo)(nil)

// NewRepo .
func NewRepo() *Repo {
	return &Repo{
		Storage: make(map[string]*repo.HatMod),
	}
}

// FindAll .
func (r *Repo) FindAll(limit repo.Limit, offset repo.Offset) (hats []*repo.HatMod, err error) {
	// TODO: respect limit and offset
	for _, s := range r.Storage {
		hats = append(hats, s)
	}
	return hats, nil
}

// Save .
func (r *Repo) Save(hm repo.HatMod) (*repo.HatMod, error) {

	var id string
	if hm.ID == "" && hm.Version != 0 {
		return nil, repo.ErrVersionMismatch
	} else if hm.ID == "" {
		id = xid.New().String()
	} else {
		id = hm.ID
	}

	mod := &repo.HatMod{
		ID:      id,
		Color:   hm.Color,
		Style:   hm.Style,
		Inches:  hm.Inches,
		Version: hm.Version + 1,
	}

	r.Storage[id] = mod
	return mod, nil
}

// Exists .
func (r *Repo) Exists(id string) (bool, error) {
	_, ok := r.Storage[id]
	return ok, nil
}

// Delete .
func (r *Repo) Delete(id string, version int) error {
	v, ok := r.Storage[id]

	if !ok {
		return repo.ErrNotFound
	}

	if v.Version != version {
		return repo.ErrVersionMismatch
	}

	delete(r.Storage, id)
	return nil
}

// Find .
func (r *Repo) Find(id string) (*repo.HatMod, error) {
	v, ok := r.Storage[id]
	if ok {
		return v, nil
	}
	return nil, nil
}

// Multi .
func (r *Repo) Multi() error {
	logrus.Debug("inmem.Multi()")
	return nil
}

// Exec .
func (r *Repo) Exec() error {
	logrus.Debug("inmem.Exec()")
	return nil
}

// Discard .
func (r *Repo) Discard() error {
	logrus.Debug("inmem.Discard()")
	return nil
}

////////// connection related

// OpenConn .
func (r *Repo) OpenConn() error {
	logrus.Debug("inmem.OpenConn()")
	return nil
}

// CloseConn implements HatRepo.CloseConn()
func (r *Repo) CloseConn() error {
	logrus.Debug("inmem.CloseConn()")
	return nil
}

// Close implements HatRepo.Close()
func (r *Repo) Close() error {
	logrus.Debug("inmem.Close()")
	r.Storage = nil
	return nil
}

// Clone implements HatRepo.Clone()
func (r *Repo) Clone() (repo.HatRepo, error) {
	logrus.Debug("inmem.Clone()")
	return &Repo{
		Storage: r.Storage,
	}, nil
}
